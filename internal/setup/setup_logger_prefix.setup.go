package setup

import (
	"context"
	"go.opentelemetry.io/otel/trace"
	"os"

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
	var kvs = []interface{}{
		"app",
		s.LoggerPrefixField().String(),
	}
	kvs = append(kvs, "tracer", s.withLoggerTracer())
	return log.With(logger, kvs...)
}

// withLoggerTracer returns a traceid valuer.
func (s *engines) withLoggerTracer() log.Valuer {
	return func(ctx context.Context) interface{} {
		var res string
		if span := trace.SpanContextFromContext(ctx); span.HasTraceID() {
			res += `traceId="` + span.TraceID().String() + `"`
		}
		if span := trace.SpanContextFromContext(ctx); span.HasSpanID() {
			res += ` spanId="` + span.SpanID().String() + `"`
		}
		return res
	}
}
