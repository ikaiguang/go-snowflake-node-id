package services

import (
	"context"
	commonv1 "github.com/ikaiguang/go-snowflake-node-id/api/common/v1"
	snowflakeerrorv1 "github.com/ikaiguang/go-snowflake-node-id/api/snowflake-service/v1/errors"
	snowflakev1 "github.com/ikaiguang/go-snowflake-node-id/api/snowflake-service/v1/resources"
	snowflakeservicev1 "github.com/ikaiguang/go-snowflake-node-id/api/snowflake-service/v1/services"
	snowflakeutil "github.com/ikaiguang/go-snowflake-node-id/business-util/snowflake"
	errorutil "github.com/ikaiguang/go-srv-kit/error"
	"strings"
)

// worker ...
type worker struct {
	snowflakeservicev1.UnimplementedSrvSnowflakeWorkerServer

	locker     snowflakeutil.Locker
	workerRepo snowflakeutil.WorkerRepo
}

// NewWorker ...
func NewWorker(
	locker snowflakeutil.Locker,
	workerRepo snowflakeutil.WorkerRepo,
) snowflakeservicev1.SrvSnowflakeWorkerServer {
	return &worker{
		locker:     locker,
		workerRepo: workerRepo,
	}
}

// GetNodeId 获取节点ID
func (s *worker) GetNodeId(ctx context.Context, in *snowflakev1.GetNodeIdReq) (*snowflakev1.SnowflakeWorker, error) {
	in.InstanceId = strings.TrimSpace(in.InstanceId)
	if in.InstanceId == "" {
		reason := snowflakeerrorv1.ERROR_WORKER_INSTANCE_ID_EMPTY.String()
		message := "实例ID不能为空"
		err := errorutil.NotFound(reason, message)
		return nil, err
	}

	// 锁
	unlocker, err := s.locker.Lock(ctx, in.InstanceId)
	if err != nil {
		reason := commonv1.ERROR_INTERNAL_SERVER.String()
		message := "服务器错误"
		err = errorutil.InternalServer(reason, message)
		return nil, err
	}
	defer func() { _, _ = unlocker.Unlock(ctx) }()

	return s.workerRepo.GetNodeId(ctx, in)
}

// ExtendNodeId 续期
func (s *worker) ExtendNodeId(ctx context.Context, in *snowflakev1.ExtendNodeIdReq) (*snowflakev1.Result, error) {
	in.InstanceId = strings.TrimSpace(in.InstanceId)
	if in.InstanceId == "" {
		reason := snowflakeerrorv1.ERROR_WORKER_INSTANCE_ID_EMPTY.String()
		message := "实例ID不能为空"
		err := errorutil.NotFound(reason, message)
		return nil, err
	}
	return s.workerRepo.ExtendNodeId(ctx, in)
}
