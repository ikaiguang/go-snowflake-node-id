package pingroute

import (
	services "github.com/ikaiguang/go-snowflake-node-id/internal/application/ping/service"
	"github.com/ikaiguang/go-snowflake-node-id/internal/setup"
	logutil "github.com/ikaiguang/go-srv-kit/log"
	stdlog "log"

	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	pingservicev1 "github.com/ikaiguang/go-srv-kit/api/ping/v1/services"
)

// RegisterPingRoutes 注册路由
func RegisterPingRoutes(engineHandler setup.Engine, hs *http.Server, gs *grpc.Server) {
	// 日志
	logger, _, err := engineHandler.Logger()
	if err != nil {
		logutil.Fatal(err)
		return
	}

	ping := services.NewPingService(logger)
	stdlog.Println("|*** 注册路由：NewPingService")
	pingservicev1.RegisterSrvPingHTTPServer(hs, ping)
	pingservicev1.RegisterSrvPingServer(gs, ping)
}
