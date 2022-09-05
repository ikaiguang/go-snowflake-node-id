package workersrv

import (
	"context"
	errorv1 "github.com/ikaiguang/go-srv-kit/api/error/v1"
	errorutil "github.com/ikaiguang/go-srv-kit/error"
	"strings"

	apiv1 "github.com/ikaiguang/go-snowflake-node-id/api"
	servicev1 "github.com/ikaiguang/go-snowflake-node-id/api/node-id/v1/services"
	nodeid "github.com/ikaiguang/go-snowflake-node-id/node-id"
)

// worker ...
type worker struct {
	servicev1.UnimplementedSrvWorkerServer

	locker     nodeid.Locker
	workerRepo nodeid.WorkerRepo
}

// NewWorker ...
func NewWorker(
	locker nodeid.Locker,
	workerRepo nodeid.WorkerRepo,
) servicev1.SrvWorkerServer {
	return &worker{
		locker:     locker,
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

	// 锁
	unlocker, err := s.locker.Lock(ctx, in.InstanceId)
	if err != nil {
		reason := errorv1.ERROR_INTERNAL_SERVER.String()
		message := "服务器错误"
		err = errorutil.InternalServer(reason, message)
		return nil, err
	}
	defer func() { _, _ = unlocker.Unlock(ctx) }()

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
