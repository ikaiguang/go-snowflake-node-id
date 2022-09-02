package workersrv

import (
	"context"
	errorutil "github.com/ikaiguang/go-srv-kit/error"
	"github.com/patrickmn/go-cache"
	"strings"
	"sync"
	"time"

	apiv1 "github.com/ikaiguang/go-snowflake-node-id/api"
	servicev1 "github.com/ikaiguang/go-snowflake-node-id/api/node-id/v1/services"
	nodeid "github.com/ikaiguang/go-snowflake-node-id/node-id"
)

// worker ...
type worker struct {
	servicev1.UnimplementedSrvWorkerServer

	cacheHandler *cache.Cache
	workerRepo   nodeid.WorkerRepo
}

// NewWorker ...
func NewWorker(
	cacheHandler *cache.Cache,
	workerRepo nodeid.WorkerRepo,
) servicev1.SrvWorkerServer {
	return &worker{
		cacheHandler: cacheHandler,
		workerRepo:   workerRepo,
	}
}

// getOrSetInstanceLocker 获取或设置 实例锁
// cache 已有读写锁，无需在添加额外的锁
func (s *worker) getOrSetInstanceLocker(ctx context.Context, instanceID string) *sync.Mutex {
	// 添加锁
	locker, ok := s.cacheHandler.Get(instanceID)
	if ok {
		return locker.(*sync.Mutex)
	}

	// 设置锁
	mu := &sync.Mutex{}
	s.cacheHandler.Set(instanceID, mu, time.Minute*5)
	return mu
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

	locker := s.getOrSetInstanceLocker(ctx, in.InstanceId)
	locker.Lock()
	defer locker.Unlock()

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
