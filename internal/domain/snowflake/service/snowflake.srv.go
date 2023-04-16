package srvs

import (
	"context"
	"encoding/json"
	commonv1 "github.com/ikaiguang/go-snowflake-node-id/api/common/v1"
	snowflakeerrorv1 "github.com/ikaiguang/go-snowflake-node-id/api/snowflake-service/v1/errors"
	snowflakev1 "github.com/ikaiguang/go-snowflake-node-id/api/snowflake-service/v1/resources"
	snowflakeutil "github.com/ikaiguang/go-snowflake-node-id/business-util/snowflake"
	entities "github.com/ikaiguang/go-snowflake-node-id/internal/domain/snowflake/entity"
	repos "github.com/ikaiguang/go-snowflake-node-id/internal/domain/snowflake/repo"
	errorutil "github.com/ikaiguang/go-srv-kit/error"
	"time"
)

const (
	DefaultMaxNodeId      = 1023             // 最大节点ID
	DefaultIdleDuration   = 16 * time.Second // 空闲ID时间：超过16s不续期，节点ID变为空闲的ID
	DefaultExtentDuration = 5 * time.Second  // 续期间隔时间
)

// Options 配置
type Options struct {
	// MaxNodeID 节点ID；默认 DefaultMaxNodeId = 1023
	MaxNodeID int64
	// IdleDuration 空闲ID时间：超过此时间不续期，节点ID变为空闲的ID；
	// 默认 DefaultIdleDuration = 16s
	IdleDuration time.Duration
}

// SnowflakeSrv ...
type SnowflakeSrv struct {
	opts                *Options
	locker              snowflakeutil.Locker
	snowflakeWorkerRepo repos.SnowflakeWorkerRepo
}

// NewSnowflakeSrv ...
func NewSnowflakeSrv(
	opts Options,
	locker snowflakeutil.Locker,
	snowflakeWorkerRepo repos.SnowflakeWorkerRepo,
) *SnowflakeSrv {
	// maxNodeID & idleDuration
	if opts.MaxNodeID < 1 {
		opts.MaxNodeID = DefaultMaxNodeId
	}
	if opts.IdleDuration < 1 {
		opts.IdleDuration = DefaultIdleDuration
	}
	return &SnowflakeSrv{
		opts:                &opts,
		locker:              locker,
		snowflakeWorkerRepo: snowflakeWorkerRepo,
	}
}

// initGetNodeIdRequest 默认值
func (s *SnowflakeSrv) initGetNodeIdRequest(req *snowflakev1.GetNodeIdReq) {
	if len(req.Endpoints) == 0 {
		req.Endpoints = []string{}
	}
	if req.Metadata == nil {
		req.Metadata = map[string]string{}
	}
}

// GetNodeId 获取节点ID
func (s *SnowflakeSrv) GetNodeId(ctx context.Context, req *snowflakev1.GetNodeIdReq) (resp *entities.SnowflakeWorker, err error) {
	// 锁
	unlocker, err := s.locker.Lock(ctx, req.InstanceId)
	if err != nil {
		reason := commonv1.ERROR_INTERNAL_SERVER.String()
		message := "服务器错误"
		err = errorutil.InternalServer(reason, message)
		return nil, err
	}
	defer func() { _, _ = unlocker.Unlock(ctx) }()

	// 默认值
	s.initGetNodeIdRequest(req)

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
func (s *SnowflakeSrv) ExtendNodeId(ctx context.Context, req *snowflakev1.ExtendNodeIdReq) (success bool, err error) {
	queryReq := &entities.SnowflakeWorker{
		Id:              req.Id,
		InstanceId:      req.InstanceId,
		SnowflakeNodeId: req.NodeId,
	}
	queryReq.NodeUuid = queryReq.GenNodeUUID()
	dataModel, isNotFound, err := s.snowflakeWorkerRepo.QueryOneByIDAndNodeUUID(ctx, queryReq)
	if err != nil {
		reason := commonv1.ERROR_INTERNAL_SERVER.String()
		message := "服务内部错误"
		err = errorutil.InternalServer(reason, message, err)
		return success, err
	}
	if isNotFound {
		reason := snowflakeerrorv1.ERROR_CANNOT_FOUNT_EXTEND_ID.String()
		message := "未找到续期的节点ID"
		err = errorutil.NotFound(reason, message)
		return success, err
	}

	// 续期
	err = s.snowflakeWorkerRepo.ExtendNodeID(ctx, dataModel)
	if err != nil {
		reason := commonv1.ERROR_INTERNAL_SERVER.String()
		message := "服务内部错误"
		err = errorutil.InternalServer(reason, message, err)
		return success, err
	}

	success = true
	return success, err
}

// getMissingNodeID 获取缺失的ID
func (s *SnowflakeSrv) getMissingNodeID(ctx context.Context, req *snowflakev1.GetNodeIdReq) (resp *snowflakev1.SnowflakeWorker, hasValidID bool, err error) {
	// todo 未实现
	return resp, hasValidID, err
}

// getIdleNodeID 获取闲置的ID
func (s *SnowflakeSrv) getIdleNodeID(ctx context.Context, req *snowflakev1.GetNodeIdReq) (resp *entities.SnowflakeWorker, hasValidID bool, err error) {
	// 获取有效的ID
	idleReq := &entities.InstanceIdleNodeIDReq{
		InstanceId:            req.InstanceId,
		MaxInstanceExtendTime: time.Now().Add(-s.opts.IdleDuration),
	}

	nodeIDModel, isNotFound, err := s.snowflakeWorkerRepo.QueryIdleNodeIDByInstanceID(ctx, idleReq)
	if err != nil {
		reason := commonv1.ERROR_INTERNAL_SERVER.String()
		message := "服务内部错误"
		err = errorutil.InternalServer(reason, message, err)
		return resp, hasValidID, err
	}
	if isNotFound {
		return resp, hasValidID, err
	}
	if nodeIDModel.SnowflakeNodeId > s.opts.MaxNodeID {
		return resp, hasValidID, err
	}

	// 查询
	newDataModel := s.assembleSnowflakeWorkerFromNodeRequest(req, nodeIDModel.SnowflakeNodeId)
	oldDataModel, isNotFound, err := s.snowflakeWorkerRepo.QueryOneByNodeUUID(ctx, newDataModel.NodeUuid)
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
	err = s.snowflakeWorkerRepo.Update(ctx, newDataModel)
	if err != nil {
		reason := commonv1.ERROR_INTERNAL_SERVER.String()
		message := "服务内部错误"
		err = errorutil.InternalServer(reason, message, err)
		return resp, hasValidID, err
	}

	hasValidID = true

	return newDataModel, hasValidID, err
}

// getLastNodeID 获取下一个ID
func (s *SnowflakeSrv) getLastNodeID(ctx context.Context, req *snowflakev1.GetNodeIdReq) (resp *entities.SnowflakeWorker, hasValidID bool, err error) {
	// 获取有效的ID
	ids, err := s.snowflakeWorkerRepo.QueryMaxNodeIDByInstanceID(ctx, req.InstanceId)
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
	if nodeID > s.opts.MaxNodeID {
		return resp, hasValidID, err
	}

	// 存储
	dataModel := s.assembleSnowflakeWorkerFromNodeRequest(req, nodeID)
	err = s.snowflakeWorkerRepo.Create(ctx, dataModel)
	if err != nil {
		reason := commonv1.ERROR_INTERNAL_SERVER.String()
		message := "服务内部错误"
		err = errorutil.InternalServer(reason, message, err)
		return resp, hasValidID, err
	}

	hasValidID = true

	return dataModel, hasValidID, err
}

// assembleSnowflakeWorkerFromNodeRequest 存储节点ID
func (s *SnowflakeSrv) assembleSnowflakeWorkerFromNodeRequest(req *snowflakev1.GetNodeIdReq, nodeID int64) (dataModel *entities.SnowflakeWorker) {
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
