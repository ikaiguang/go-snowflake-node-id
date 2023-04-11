package pingroute

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/mux"
	services "github.com/ikaiguang/go-snowflake-node-id/internal/application/ping/service"
	srvs "github.com/ikaiguang/go-snowflake-node-id/internal/domain/ping/service"
	"github.com/ikaiguang/go-snowflake-node-id/internal/setup"
	logutil "github.com/ikaiguang/go-srv-kit/log"
	stdlog "log"
)

// RegisterWssRoutes 注册路由
// ref https://github.com/go-kratos/examples/blob/main/ws/main.go
func RegisterWssRoutes(engineHandler setup.Engine, hs *http.Server, gs *grpc.Server) {
	// 日志
	logger, _, err := engineHandler.Logger()
	if err != nil {
		logutil.Fatal(err)
		return
	}

	wssSrv := srvs.NewWebsocketSrv(logger)
	handler := services.NewWebsocketService(wssSrv)

	router := mux.NewRouter()
	router.HandleFunc("/ws/v1/websocket", handler.TestWebsocket)

	stdlog.Println("|*** 注册路由：Websocket")
	hs.Handle("/ws/v1/websocket", router)
}
