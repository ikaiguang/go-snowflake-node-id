package setuppkg

import (
	"context"
	authutil "github.com/ikaiguang/go-snowflake-node-id/business-util/auth"
	contextutil "github.com/ikaiguang/go-srv-kit/kratos/context"
	headerutil "github.com/ikaiguang/go-srv-kit/kratos/header"
	"go.opentelemetry.io/otel/trace"
	"os"
	"strconv"

	"github.com/go-kratos/kratos/v2/log"

	iputil "github.com/ikaiguang/go-srv-kit/kit/ip"
)

// LoggerPrefixField .
func (s *engines) LoggerPrefixField() *LoggerPrefixField {
	s.loggerPrefixFieldMutex.Do(func() {
		s.loggerPrefixField = s.assemblyLoggerPrefixField()
	})
	return s.loggerPrefixField
}

// assemblyLoggerPrefixField 组装日志前缀
func (s *engines) assemblyLoggerPrefixField() *LoggerPrefixField {
	appConfig := s.AppConfig()
	if appConfig == nil {
		return &LoggerPrefixField{
			ServerIP: iputil.LocalIP(),
		}
	}

	fields := &LoggerPrefixField{
		AppName:    appConfig.Name,
		AppVersion: appConfig.Version,
		AppEnv:     appConfig.Env,
		AppBranch:  appConfig.EnvBranch,
		ServerIP:   iputil.LocalIP(),
	}
	fields.Hostname, _ = os.Hostname()
	return fields
}

// withLoggerPrefix ...
func (s *engines) withLoggerPrefix(logger log.Logger) log.Logger {
	//var kvs = []interface{}{
	//	"app",
	//	s.LoggerPrefixField().String(),
	//}
	var kvs = s.LoggerPrefixField().Prefix()
	kvs = append(kvs, "tracer", s.withLoggerTracer())
	return log.With(logger, kvs...)
}

// withLoggerTracer returns a traceid valuer.
func (s *engines) withLoggerTracer() log.Valuer {
	return func(ctx context.Context) interface{} {
		var (
			res        string
			hasTraceId bool
		)
		span := trace.SpanContextFromContext(ctx)
		if span.HasTraceID() {
			hasTraceId = true
			res += `traceId="` + span.TraceID().String() + `"`
		}
		if !hasTraceId {
			if httpTr, ok := contextutil.MatchHTTPServerContext(ctx); ok {
				if traceId := httpTr.RequestHeader().Get(headerutil.RequestID); traceId != "" {
					res += `traceId="` + traceId + `"`
				}
			} else if grpcTr, ok := contextutil.MatchGRPCServerContext(ctx); ok {
				if traceId := grpcTr.RequestHeader().Get(headerutil.RequestID); traceId != "" {
					res += `traceId="` + traceId + `"`
				}
			}
		}
		if span.HasSpanID() {
			res += ` spanId="` + span.SpanID().String() + `"`
		}
		if claims, ok := authutil.FromAuthContext(ctx); ok && claims.Payload != nil {
			if claims.Payload.Id > 0 {
				res += ` userId="` + strconv.FormatUint(claims.Payload.Id, 10) + `"`
			} else {
				res += ` userId="` + claims.Payload.Uid + `"`
			}
		}
		return res
	}
}
