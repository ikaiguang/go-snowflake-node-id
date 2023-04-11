package servers

import (
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	middlewareutil "github.com/ikaiguang/go-snowflake-node-id/business-util/middleware"
	"github.com/ikaiguang/go-snowflake-node-id/internal/setup"
	apppkg "github.com/ikaiguang/go-snowflake-node-id/pkg/app"
	stdlog "log"
)

var _ metadata.Option

// NewGRPCServer new a gRPC server.
func NewGRPCServer(engineHandler setup.Engine) (srv *grpc.Server, err error) {
	grpcConfig := engineHandler.GRPCConfig()
	stdlog.Printf("|*** 加载：GRPC服务：%s\n", grpcConfig.Addr)

	// 日志
	//logger, _, err := engineHandler.Logger()
	//if err != nil {
	//	return srv, err
	//}

	// options
	var opts []grpc.ServerOption
	if grpcConfig.Network != "" {
		opts = append(opts, grpc.Network(grpcConfig.Network))
	}
	if grpcConfig.Addr != "" {
		opts = append(opts, grpc.Address(grpcConfig.Addr))
	}
	if grpcConfig.Timeout != nil {
		opts = append(opts, grpc.Timeout(grpcConfig.Timeout.AsDuration()))
	}

	// ===== 中间件 =====
	var middlewareSlice = middlewareutil.DefaultMiddlewares()
	// tracer
	settingConfig := engineHandler.BaseSettingConfig()
	if settingConfig != nil && settingConfig.EnableServiceTracer {
		stdlog.Println("|*** 加载：服务追踪：GRPC")
		if err = middlewareutil.SetTracerProvider(engineHandler); err != nil {
			return srv, err
		}
		middlewareSlice = append(middlewareSlice, tracing.Server())
	}
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
	//stdlog.Println("|*** 加载：JWT中间件：GRPC")
	jwtMiddleware, err := middlewareutil.NewJWTMiddleware(engineHandler, getAuthWhiteList())
	if err != nil {
		return srv, err
	}
	middlewareSlice = append(middlewareSlice, jwtMiddleware)

	// 中间件选项
	opts = append(opts, grpc.Middleware(middlewareSlice...))

	// 服务
	srv = grpc.NewServer(opts...)
	//v1.RegisterGreeterServer(srv, greeter)

	return srv, err
}

// RegisterGRPCRoute 注册路由
func RegisterGRPCRoute(engineHandler setup.Engine, srv *grpc.Server) (err error) {
	stdlog.Println("|*** 注册GRPC路由：...")
	return err
}
