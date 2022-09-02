package workersrv

import (
	"context"
	errorutil "github.com/ikaiguang/go-srv-kit/error"
	"strings"

	apiv1 "github.com/ikaiguang/go-snowflake-node-id/api"
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
func (s *worker) GetNodeId(ctx context.Context, in *apiv1.GetNodeIdReq) (*apiv1.SnowflakeWorkerNode, error) {
	in.InstanceId = strings.TrimSpace(in.InstanceId)
	if in.InstanceId == "" {
		reason := apiv1.ERROR_WORKER_INSTANCE_ID_EMPTY.String()
		message := "实例ID不能为空"
		err := errorutil.NotFound(reason, message)
		return nil, err
	}
	return s.workerRepo.GetNodeId(ctx, in)
}

// ExtendNodeId 续期
func (s *worker) ExtendNodeId(ctx context.Context, in *apiv1.ExtendNodeIdReq) (*apiv1.Result, error) {
	in.InstanceId = strings.TrimSpace(in.InstanceId)
	if in.InstanceId == "" {
		reason := apiv1.ERROR_WORKER_INSTANCE_ID_EMPTY.String()
		message := "实例ID不能为空"
		err := errorutil.NotFound(reason, message)
		return nil, err
	}
	return s.workerRepo.ExtendNodeId(ctx, in)
}
