package routes

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	pingroute "github.com/ikaiguang/go-snowflake-node-id/internal/route/ping"
	snowflakeroute "github.com/ikaiguang/go-snowflake-node-id/internal/route/snowflake"
	"github.com/ikaiguang/go-snowflake-node-id/internal/setup"
	stdlog "log"
)

// RegisterRoutes 注册路由
func RegisterRoutes(engineHandler setup.Engine, hs *http.Server, gs *grpc.Server) (err error) {
	stdlog.Println("|*** 注册路由：...")

	// root
	pingroute.RegisterRootRoutes(engineHandler, hs, gs)

	// websocket
	pingroute.RegisterWssRoutes(engineHandler, hs, gs)

	// testdata
	pingroute.RegisterPingRoutes(engineHandler, hs, gs)
	pingroute.RegisterTestdataRoutes(engineHandler, hs, gs)

	// snowflake
	snowflakeroute.RegisterRoutes(engineHandler, hs, gs)

	return err
}
