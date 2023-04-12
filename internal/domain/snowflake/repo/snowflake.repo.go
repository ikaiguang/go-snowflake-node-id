// Package repos
// Code generated by ikaiguang. <https://github.com/ikaiguang>
package repos

import (
	context "context"
	entities "github.com/ikaiguang/go-snowflake-node-id/internal/domain/snowflake/entity"
)

// SnowflakeWorkerRepo repo
type SnowflakeWorkerRepo interface {
	Create(ctx context.Context, dataModel *entities.SnowflakeWorker) error
	Update(ctx context.Context, dataModel *entities.SnowflakeWorker) error
	ExtendNodeID(ctx context.Context, dataModel *entities.SnowflakeWorker) (err error)
	QueryOneByNodeUUID(ctx context.Context, nodeUUID string) (dataModel *entities.SnowflakeWorker, isNotFound bool, err error)
	QueryOneByIDAndNodeUUID(ctx context.Context, req *entities.SnowflakeWorker) (dataModel *entities.SnowflakeWorker, isNotFound bool, err error)
	QueryMaxNodeIDByInstanceID(ctx context.Context, instanceID string) (dataModels []*entities.InstanceMaxNodeID, err error)
	QueryIdleNodeIDByInstanceID(ctx context.Context, req *entities.InstanceIdleNodeIDReq) (dataModel *entities.InstanceMaxNodeID, isNotFound bool, err error)
}
