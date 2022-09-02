package workerroute

import (
	stdlog "log"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	servicev1 "github.com/ikaiguang/go-snowflake-node-id/api/node-id/v1/services"
	workersrv "github.com/ikaiguang/go-snowflake-node-id/internal/application/service/node-id"
	"github.com/ikaiguang/go-snowflake-node-id/internal/setup"
	nodeid "github.com/ikaiguang/go-snowflake-node-id/node-id"
)

// RegisterRoutes 注册路由
func RegisterRoutes(engineHandler setup.Engine, hs *http.Server, gs *grpc.Server) (err error) {
	stdlog.Println("|*** 注册路由：NodeID")

	// 数据库
	dbConn, err := engineHandler.GetMySQLGormDB()
	if err != nil {
		return err
	}

	// node-id
	workerRepo, err := nodeid.NewWorker(
		nodeid.WithDBConn(dbConn),
		nodeid.WithIdleDuration(nodeid.DefaultIdleDuration),
		nodeid.WithMaxNodeID(nodeid.DefaultMaxNodeId),
	)
	if err != nil {
		return err
	}

	// 服务
	srv := workersrv.NewWorker(workerRepo)
	servicev1.RegisterSrvWorkerHTTPServer(hs, srv)
	servicev1.RegisterSrvWorkerServer(gs, srv)

	return err
}