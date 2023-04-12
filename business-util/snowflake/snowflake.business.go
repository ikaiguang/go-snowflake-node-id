package snowflakeutil

import (
	"context"
	"encoding/json"
	stderrors "errors"
	commonv1 "github.com/ikaiguang/go-snowflake-node-id/api/common/v1"
	snowflakeerrorv1 "github.com/ikaiguang/go-snowflake-node-id/api/snowflake-service/v1/errors"
	snowflakev1 "github.com/ikaiguang/go-snowflake-node-id/api/snowflake-service/v1/resources"
	assemblers "github.com/ikaiguang/go-snowflake-node-id/internal/application/snowflake/assembler"
	entities "github.com/ikaiguang/go-snowflake-node-id/internal/domain/snowflake/entity"
	repos "github.com/ikaiguang/go-snowflake-node-id/internal/domain/snowflake/repo"
	datas "github.com/ikaiguang/go-snowflake-node-id/internal/infra/snowflake/data"
	errorutil "github.com/ikaiguang/go-srv-kit/error"
	"time"
)

const (
	DefaultMaxNodeId      = 1023             // 最大节点ID
	DefaultIdleDuration   = 16 * time.Second // 空闲ID时间：超过16s不续期，节点ID变为空闲的ID
	DefaultExtentDuration = 5 * time.Second  // 续期间隔时间
)

// worker ...
type worker struct {
	nodeRepo repos.SnowflakeWorkerRepo

	opt *options
}

// NewWorker ...
func NewWorker(opts ...Option) (WorkerRepo, error) {
	options := &options{}
	for i := range opts {
		opts[i](options)
	}
	// db
	if options.dbConn == nil {
		err := stderrors.New("[nodeid.NewWorker] 缺少参数：dbConn")
		return nil, err
	}

	// maxNodeID & idleDuration
	if options.maxNodeID < 1 {
		options.maxNodeID = DefaultMaxNodeId
	}
	if options.idleDuration < 1 {
		options.idleDuration = DefaultIdleDuration
	}
	w := &worker{
		opt:      options,
		nodeRepo: datas.NewSnowflakeWorkerRepo(options.dbConn),
	}
	return w, nil
}

// GetNodeId 获取节点ID
func (s *worker) GetNodeId(ctx context.Context, req *snowflakev1.GetNodeIdReq) (resp *snowflakev1.SnowflakeWorker, err error) {
	if len(req.Endpoints) == 0 {
		req.Endpoints = []string{}
	}
	if req.Metadata == nil {
		req.Metadata = map[string]string{}
	}

	// 获取下一个ID
	resp, hasValidID, err := s.getLastNodeID(ctx, req)
	if err != nil {
		return resp, err
	}
	if hasValidID {
		return resp, err
	}

	// 获取闲置的ID
	resp, hasValidID, err = s.getIdleNodeID(ctx, req)
	if err != nil {
		return resp, err
	}
	if hasValidID {
		return resp, err
	}

	// 获取缺失的ID
	// 只有人为删除数据库数据的情况，才需补充此步骤
	//resp, hasValidID, err = s.getMissingNodeID(ctx, req)
	//if err != nil {
	//	return resp, err
	//}
	//if hasValidID {
	//	return resp, err
	//}

	reason := snowflakeerrorv1.ERROR_CANNOT_FOUNT_USABLE_ID
	message := "未找到可用的节点ID"
	err = errorutil.NotFound(reason.String(), message)
	return resp, err
}

// ExtendNodeId 续期
func (s *worker) ExtendNodeId(ctx context.Context, req *snowflakev1.ExtendNodeIdReq) (resp *snowflakev1.Result, err error) {
	queryReq := &entities.SnowflakeWorker{
		Id:              req.Id,
		InstanceId:      req.InstanceId,
		SnowflakeNodeId: req.NodeId,
	}
	queryReq.NodeUuid = queryReq.GenNodeUUID()
	dataModel, isNotFound, err := s.nodeRepo.QueryOneByIDAndNodeUUID(ctx, queryReq)
	if err != nil {
		reason := commonv1.ERROR_INTERNAL_SERVER.String()
		message := "服务内部错误"
		err = errorutil.InternalServer(reason, message, err)
		return resp, err
	}
	if isNotFound {
		reason := snowflakeerrorv1.ERROR_CANNOT_FOUNT_EXTEND_ID.String()
		message := "未找到续期的节点ID"
		err = errorutil.NotFound(reason, message)
		return resp, err
	}

	// 续期
	err = s.nodeRepo.ExtendNodeID(ctx, dataModel)
	if err != nil {
		reason := commonv1.ERROR_INTERNAL_SERVER.String()
		message := "服务内部错误"
		err = errorutil.InternalServer(reason, message, err)
		return resp, err
	}

	resp = &snowflakev1.Result{Success: true}
	return resp, err
}

// getMissingNodeID 获取缺失的ID
func (s *worker) getMissingNodeID(ctx context.Context, req *snowflakev1.GetNodeIdReq) (resp *snowflakev1.SnowflakeWorker, hasValidID bool, err error) {
	// todo 未实现
	return resp, hasValidID, err
}

// getIdleNodeID 获取闲置的ID
func (s *worker) getIdleNodeID(ctx context.Context, req *snowflakev1.GetNodeIdReq) (resp *snowflakev1.SnowflakeWorker, hasValidID bool, err error) {
	// 获取有效的ID
	idleReq := &entities.InstanceIdleNodeIDReq{
		InstanceId:            req.InstanceId,
		MaxInstanceExtendTime: time.Now().Add(-s.opt.idleDuration),
	}

	nodeIDModel, isNotFound, err := s.nodeRepo.QueryIdleNodeIDByInstanceID(ctx, idleReq)
	if err != nil {
		reason := commonv1.ERROR_INTERNAL_SERVER.String()
		message := "服务内部错误"
		err = errorutil.InternalServer(reason, message, err)
		return resp, hasValidID, err
	}
	if isNotFound {
		return resp, hasValidID, err
	}
	if nodeIDModel.SnowflakeNodeId > s.opt.maxNodeID {
		return resp, hasValidID, err
	}

	// 查询
	newDataModel := s.assembleNodeId(ctx, req, nodeIDModel.SnowflakeNodeId)
	oldDataModel, isNotFound, err := s.nodeRepo.QueryOneByNodeUUID(ctx, newDataModel.NodeUuid)
	if err != nil {
		reason := commonv1.ERROR_INTERNAL_SERVER.String()
		message := "服务内部错误"
		err = errorutil.InternalServer(reason, message, err)
		return resp, hasValidID, err
	}
	if isNotFound {
		reason := commonv1.ERROR_INTERNAL_SERVER.String()
		message := "查无此记录：" + newDataModel.NodeUuid
		err = errorutil.InternalServer(reason, message)
		return resp, hasValidID, err
	}
	newDataModel.Id = oldDataModel.Id

	// 存储
	err = s.nodeRepo.Update(ctx, newDataModel)
	if err != nil {
		reason := commonv1.ERROR_INTERNAL_SERVER.String()
		message := "服务内部错误"
		err = errorutil.InternalServer(reason, message, err)
		return resp, hasValidID, err
	}

	hasValidID = true
	resp = assemblers.AssembleSnowflakeWorker(newDataModel)
	return resp, hasValidID, err
}

// getLastNodeID 获取下一个ID
func (s *worker) getLastNodeID(ctx context.Context, req *snowflakev1.GetNodeIdReq) (resp *snowflakev1.SnowflakeWorker, hasValidID bool, err error) {
	// 获取有效的ID
	ids, err := s.nodeRepo.QueryMaxNodeIDByInstanceID(ctx, req.InstanceId)
	if err != nil {
		reason := commonv1.ERROR_INTERNAL_SERVER.String()
		message := "服务内部错误"
		err = errorutil.InternalServer(reason, message, err)
		return resp, hasValidID, err
	}
	if len(ids) != 1 {
		return resp, hasValidID, err
	}
	nodeID := ids[0].SnowflakeNodeId + 1
	if nodeID > s.opt.maxNodeID {
		return resp, hasValidID, err
	}

	// 存储
	dataModel := s.assembleNodeId(ctx, req, nodeID)
	err = s.nodeRepo.Create(ctx, dataModel)
	if err != nil {
		reason := commonv1.ERROR_INTERNAL_SERVER.String()
		message := "服务内部错误"
		err = errorutil.InternalServer(reason, message, err)
		return resp, hasValidID, err
	}

	hasValidID = true
	resp = assemblers.AssembleSnowflakeWorker(dataModel)
	return resp, hasValidID, err
}

// assembleNodeId 存储节点ID
func (s *worker) assembleNodeId(ctx context.Context, req *snowflakev1.GetNodeIdReq, nodeID int64) (dataModel *entities.SnowflakeWorker) {
	now := time.Now()
	dataModel = &entities.SnowflakeWorker{
		InstanceLaunchTime:   now,
		InstanceExtendTime:   now,
		InstanceId:           req.InstanceId,
		SnowflakeNodeId:      nodeID,
		InstanceName:         req.InstanceName,
		InstanceEndpointList: "",
		InstanceMetadata:     "",
		CreatedTime:          now,
	}
	dataModel.NodeUuid = dataModel.GenNodeUUID()
	endpoints, _ := json.Marshal(req.Endpoints)
	dataModel.InstanceEndpointList = string(endpoints)
	metadata, _ := json.Marshal(req.Metadata)
	dataModel.InstanceMetadata = string(metadata)

	return dataModel
}
