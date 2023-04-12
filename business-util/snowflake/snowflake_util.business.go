package snowflakeutil

import (
	"context"
	snowflakev1 "github.com/ikaiguang/go-snowflake-node-id/api/snowflake-service/v1/resources"
	confv1 "github.com/ikaiguang/go-srv-kit/api/conf/v1"
	gormutil "github.com/ikaiguang/go-srv-kit/data/gorm"
	mysqlutil "github.com/ikaiguang/go-srv-kit/data/mysql"
	"gorm.io/gorm"
)

// WorkerRepo ...
type WorkerRepo interface {
	// GetNodeId 获取节点ID
	GetNodeId(ctx context.Context, req *snowflakev1.GetNodeIdReq) (resp *snowflakev1.SnowflakeWorker, err error)
	// ExtendNodeId 续期
	ExtendNodeId(ctx context.Context, req *snowflakev1.ExtendNodeIdReq) (resp *snowflakev1.Result, err error)
}

// NewMysqlDB ...
func NewMysqlDB(conf *confv1.Data_MySQL, opts ...gormutil.Option) (db *gorm.DB, err error) {
	return mysqlutil.NewMysqlDB(conf, opts...)
}
