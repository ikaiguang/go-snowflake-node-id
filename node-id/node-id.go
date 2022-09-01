package nodeid

import (
	"context"
	"encoding/json"
	stderrors "errors"
	errorv1 "github.com/ikaiguang/go-srv-kit/api/error/v1"
	errorutil "github.com/ikaiguang/go-srv-kit/error"
	"strings"
	"time"

	apiv1 "github.com/ikaiguang/go-snowflake-node-id/api"
	assemblers "github.com/ikaiguang/go-snowflake-node-id/node-id/assembler"
	datas "github.com/ikaiguang/go-snowflake-node-id/node-id/data"
	entities "github.com/ikaiguang/go-snowflake-node-id/node-id/entity"
	repos "github.com/ikaiguang/go-snowflake-node-id/node-id/repo"
)

const (
	_maxNodeId = 1023
)

// worker ...
type worker struct {
	nodeRepo repos.SnowflakeWorkerNodeRepo

	// maxNodeID 节点ID；默认 _maxNodeId = 1023
	maxNodeID int64
}

// NewWorker ...
func NewWorker(opts ...Option) (*worker, error) {
	options := &options{
		maxNodeID: _maxNodeId,
	}
	for i := range opts {
		opts[i](options)
	}
	if options.dbConn == nil {
		err := stderrors.New("[nodeid.NewWorker] 缺少参数：dbConn")
		return nil, err
	}
	w := &worker{
		maxNodeID: options.maxNodeID,
		nodeRepo:  datas.NewSnowflakeWorkerNodeRepo(options.dbConn),
	}
	return w, nil
}

// GetNodeId 获取节点ID
func (s *worker) GetNodeId(ctx context.Context, req *apiv1.GetNodeIdReq) (resp *apiv1.SnowflakeWorkerNode, err error) {
	req.Id = strings.TrimSpace(req.Id)
	if len(req.Endpoints) == 0 {
		req.Endpoints = []string{}
	}
	if req.Metadata == nil {
		req.Metadata = map[string]string{}
	}

	// 获取有效的ID
	ids, err := s.nodeRepo.QueryMaxNodeIDByInstanceID(ctx, req.Id)
	if err != nil {
		reason := errorv1.ERROR_INTERNAL_SERVER.String()
		message := "服务内部错误"
		err = errorutil.InternalServer(reason, message, err)
		return resp, err
	}
	if len(ids) == 1 && ids[0].SnowflakeNodeId+1 <= s.maxNodeID {
		nodeID := ids[0].SnowflakeNodeId + 1
		dataModel, err := s.storeNodeId(ctx, req, nodeID)
		if err != nil {
			reason := errorv1.ERROR_INTERNAL_SERVER.String()
			message := "服务内部错误"
			err = errorutil.InternalServer(reason, message, err)
			return resp, err
		}
		resp = assemblers.AssembleSnowflakeWorkerNode(dataModel)
		return resp, err
	}

	return resp, err
}

// storeNodeId 存储节点ID
func (s *worker) storeNodeId(ctx context.Context, req *apiv1.GetNodeIdReq, nodeID int64) (resp *entities.SnowflakeWorkerNode, err error) {
	now := time.Now()
	resp = &entities.SnowflakeWorkerNode{
		InstanceLaunchTime:   now,
		InstanceExtendTime:   now,
		InstanceId:           req.Id,
		SnowflakeNodeId:      nodeID,
		InstanceName:         req.Name,
		InstanceEndpointList: "",
		InstanceMetadata:     "",
		CreatedTime:          now,
	}
	resp.NodeUuid = resp.GenNodeUUID()
	endpoints, _ := json.Marshal(req.Endpoints)
	resp.InstanceEndpointList = string(endpoints)
	metadata, _ := json.Marshal(req.Metadata)
	resp.InstanceMetadata = string(metadata)

	// 存储
	err = s.nodeRepo.StoreNodeID(ctx, resp)
	if err != nil {
		reason := errorv1.ERROR_INTERNAL_SERVER.String()
		message := "服务内部错误"
		err = errorutil.InternalServer(reason, message, err)
		return resp, err
	}
	return resp, err
}
