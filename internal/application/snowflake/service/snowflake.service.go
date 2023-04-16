package services

import (
	"context"
	snowflakeerrorv1 "github.com/ikaiguang/go-snowflake-node-id/api/snowflake-service/v1/errors"
	snowflakev1 "github.com/ikaiguang/go-snowflake-node-id/api/snowflake-service/v1/resources"
	snowflakeservicev1 "github.com/ikaiguang/go-snowflake-node-id/api/snowflake-service/v1/services"
	assemblers "github.com/ikaiguang/go-snowflake-node-id/internal/application/snowflake/assembler"
	srvs "github.com/ikaiguang/go-snowflake-node-id/internal/domain/snowflake/service"
	errorutil "github.com/ikaiguang/go-srv-kit/error"
	"strings"
)

// worker ...
type worker struct {
	snowflakeservicev1.UnimplementedSrvSnowflakeWorkerServer

	assembler    *assemblers.Assembler
	snowflakeSrv *srvs.SnowflakeSrv
}

// NewWorker ...
func NewWorker(
	assembler *assemblers.Assembler,
	snowflakeSrv *srvs.SnowflakeSrv,
) snowflakeservicev1.SrvSnowflakeWorkerServer {
	return &worker{
		assembler:    assembler,
		snowflakeSrv: snowflakeSrv,
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

	dataModel, err := s.snowflakeSrv.GetNodeId(ctx, in)
	if err != nil {
		return nil, err
	}
	return s.assembler.SnowflakeWorker(dataModel), nil
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
	success, err := s.snowflakeSrv.ExtendNodeId(ctx, in)
	if err != nil {
		return nil, err
	}
	return &snowflakev1.Result{Success: success}, nil
}
