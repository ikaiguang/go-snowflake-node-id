package pingroute

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/mux"
	"github.com/ikaiguang/go-snowflake-node-id/internal/setup"
	apppkg "github.com/ikaiguang/go-snowflake-node-id/pkg/app"
	pingv1 "github.com/ikaiguang/go-srv-kit/api/ping/v1/resources"
	stdlog "log"
)

// RegisterRootRoutes 注册路由
func RegisterRootRoutes(engineHandler setup.Engine, hs *http.Server, gs *grpc.Server) {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := &pingv1.PingResp{
			Message: "Hello World!",
		}
		err := apppkg.ResponseEncoder(w, r, data)
		if err != nil {
			apppkg.ErrorEncoder(w, r, err)
		}
	})

	stdlog.Println("|*** 注册路由：Root(/)")
	hs.Handle("/", router)
}
