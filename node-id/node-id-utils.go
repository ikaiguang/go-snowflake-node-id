package nodeid

import (
	"context"
	confv1 "github.com/ikaiguang/go-srv-kit/api/conf/v1"
	gormutil "github.com/ikaiguang/go-srv-kit/data/gorm"
	mysqlutil "github.com/ikaiguang/go-srv-kit/data/mysql"
	"gorm.io/gorm"

	apiv1 "github.com/ikaiguang/go-snowflake-node-id/api"
)

// WorkerRepo ...
type WorkerRepo interface {
	// GetNodeId 获取节点ID
	GetNodeId(ctx context.Context, req *apiv1.GetNodeIdReq) (resp *apiv1.SnowflakeWorkerNode, err error)
	// ExtendNodeId 续期
	ExtendNodeId(ctx context.Context, req *apiv1.ExtendNodeIdReq) (resp *apiv1.Result, err error)
}

// NewMysqlDB ...
func NewMysqlDB(conf *confv1.Data_MySQL, opts ...gormutil.Option) (db *gorm.DB, err error) {
	return mysqlutil.NewMysqlDB(conf, opts...)
}
