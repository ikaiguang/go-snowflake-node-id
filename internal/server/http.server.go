package servers

import (
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/http"
	middlewareutil "github.com/ikaiguang/go-snowflake-node-id/business-util/middleware"
	"github.com/ikaiguang/go-snowflake-node-id/internal/setup"
	apppkg "github.com/ikaiguang/go-snowflake-node-id/pkg/app"
	stdlog "log"
)

var _ metadata.Option

// NewHTTPServer new HTTP server.
func NewHTTPServer(engineHandler setup.Engine) (srv *http.Server, err error) {
	httpConfig := engineHandler.HTTPConfig()
	stdlog.Printf("|*** 加载：HTTP服务：%s\n", httpConfig.Addr)

	// 日志
	//logger, _, err := engineHandler.Logger()
	//if err != nil {
	//	return srv, err
	//}

	// options
	var opts []http.ServerOption
	//var opts = []http.ServerOption{
	//	http.Filter(middlewareutil.NewCORS()),
	//}
	if httpConfig.Network != "" {
		opts = append(opts, http.Network(httpConfig.Network))
	}
	if httpConfig.Addr != "" {
		opts = append(opts, http.Address(httpConfig.Addr))
	}
	if httpConfig.Timeout != nil {
		opts = append(opts, http.Timeout(httpConfig.Timeout.AsDuration()))
	}

	// 响应
	opts = append(opts, http.ResponseEncoder(apppkg.ResponseEncoder))
	opts = append(opts, http.ErrorEncoder(apppkg.ErrorEncoder))

	// ===== 中间件 =====
	var middlewareSlice = middlewareutil.DefaultMiddlewares()
	// tracer
	settingConfig := engineHandler.BaseSettingConfig()
	if settingConfig != nil && settingConfig.EnableServiceTracer {
		stdlog.Println("|*** 加载：服务追踪：HTTP")
		if err = middlewareutil.SetTracerProvider(engineHandler); err != nil {
			return srv, err
		}
		middlewareSlice = append(middlewareSlice, tracing.Server())
	}
	// 请求头
	middlewareSlice = append(middlewareSlice, middlewareutil.RequestHeader())
	// 中间件日志
	middleLogger, _, err := engineHandler.LoggerMiddleware()
	if err != nil {
		return srv, err
	}
	// 日志输出
	//errorutil.DefaultStackTracerDepth += 2
	middlewareSlice = append(middlewareSlice, apppkg.ServerLog(
		middleLogger,
		//middlewareutil.WithDefaultSkip(),
	))
	// jwt
	//stdlog.Println("|*** 加载：JWT中间件：HTTP")
	jwtMiddleware, err := middlewareutil.NewJWTMiddleware(engineHandler, getAuthWhiteList())
	if err != nil {
		return srv, err
	}
	middlewareSlice = append(middlewareSlice, jwtMiddleware)

	// 中间件选项
	opts = append(opts, http.Middleware(middlewareSlice...))

	// 服务
	srv = http.NewServer(opts...)
	//v1.RegisterGreeterHTTPServer(srv, greeter)

	return srv, err
}

// RegisterHTTPRoute 注册路由
func RegisterHTTPRoute(engineHandler setup.Engine, srv *http.Server) (err error) {
	stdlog.Println("|*** 注册HTTP路由：...")
	return err
}
