package setuppkg

import (
	tokenutil "github.com/ikaiguang/go-snowflake-node-id/business-util/token"
	registrypkg "github.com/ikaiguang/go-snowflake-node-id/pkg/registry"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"io"
	"sync"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/hashicorp/consul/api"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// engines 引擎模块
type engines struct {
	Config

	// registryType 服务注册类型
	registryType registrypkg.RegistryType

	// loggerPrefixFieldMutex 日志前缀
	loggerPrefixFieldMutex sync.Once
	loggerPrefixField      *LoggerPrefixField

	// loggerFileWriterMutex 日志文件写手柄
	loggerFileWriterMutex sync.Once
	loggerFileWriter      io.Writer

	// debugHelperCloseFnSlice debug工具
	debugHelperCloseFnSlice []io.Closer

	// loggerMutex 日志
	loggerMutex                  sync.Once
	logger                       log.Logger
	loggerCloseFnSlice           []io.Closer
	loggerHelperMutex            sync.Once
	loggerHelper                 log.Logger
	loggerHelperCloseFnSlice     []io.Closer
	loggerMiddlewareMutex        sync.Once
	loggerMiddleware             log.Logger
	loggerMiddlewareCloseFnSlice []io.Closer

	// mysqlGormMutex mysql gorm
	mysqlGormMutex sync.Once
	mysqlGormDB    *gorm.DB

	// postgresGormMutex mysql gorm
	postgresGormMutex sync.Once
	postgresGormDB    *gorm.DB

	// redisClientMutex redis客户端
	redisClientMutex sync.Once
	redisClient      *redis.Client

	// consulClientMutex consul客户端
	consulClientMutex sync.Once
	consulClient      *api.Client

	// jaegerTraceExporterMutex jaeger trace
	jaegerTraceExporterMutex sync.Once
	jaegerTraceExporter      *jaeger.Exporter

	// snowflakeStopChannel 雪花算法
	snowflakeStopChannel chan int

	// authTokenRepoMutex 验证Token工具
	authTokenRepoMutex sync.Once
	authTokenRepo      tokenutil.AuthTokenRepo
}

// initEngine ...
func initEngine(conf Config) *engines {
	return &engines{
		Config: conf,
	}
}
