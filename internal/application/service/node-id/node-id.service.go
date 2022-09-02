package workersrv

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	api "github.com/ikaiguang/go-snowflake-node-id/api"
	servicev1 "github.com/ikaiguang/go-snowflake-node-id/api/node-id/v1/services"
	nodeid "github.com/ikaiguang/go-snowflake-node-id/node-id"
)

// worker ...
type worker struct {
	servicev1.UnimplementedSrvWorkerServer

	workerRepo nodeid.WorkerRepo
}

// NewWorker ...
func NewWorker(workerRepo nodeid.WorkerRepo) servicev1.SrvWorkerServer {
	return &worker{
		workerRepo: workerRepo,
	}
}

// GetNodeId 获取节点ID
func (s *worker) GetNodeId(context.Context, *api.GetNodeIdReq) (*api.SnowflakeWorkerNode, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNodeId not implemented")
}

// ExtendNodeId 续期
func (s *worker) ExtendNodeId(context.Context, *api.ExtendNodeIdReq) (*api.Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExtendNodeId not implemented")
}
