package setuppkg

import (
	strerrors "errors"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/hashicorp/consul/api"
	commonv1 "github.com/ikaiguang/go-snowflake-node-id/api/common/v1"
	tokenutil "github.com/ikaiguang/go-snowflake-node-id/business-util/token"
	configv1 "github.com/ikaiguang/go-snowflake-node-id/internal/conf"
	registrypkg "github.com/ikaiguang/go-snowflake-node-id/pkg/registry"
	confv1 "github.com/ikaiguang/go-srv-kit/api/conf/v1"
	pkgerrors "github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"gorm.io/gorm"
	"io"
	"strings"
)

var (
	_ Config = &configuration{}
	_ Engine = &engines{}

	// ErrUnimplemented 未实现
	ErrUnimplemented = strerrors.New("unimplemented")
	// ErrUninitialized 未初始化
	ErrUninitialized = strerrors.New("uninitialized")
)

// IsUnimplementedError 未实现
func IsUnimplementedError(err error) bool {
	return strerrors.Is(pkgerrors.Cause(err), ErrUnimplemented)
}

// IsUninitializedError 未初始化
func IsUninitializedError(err error) bool {
	return strerrors.Is(pkgerrors.Cause(err), ErrUninitialized)
}

// LoggerPrefixField with logger fields.
type LoggerPrefixField struct {
	AppName    string `json:"name"`
	AppVersion string `json:"version"`
	AppEnv     string `json:"env"`
	AppBranch  string `json:"envBranch"`
	Hostname   string `json:"hostname"`
	ServerIP   string `json:"serverIP"`
}

// String 日志前缀
func (s *LoggerPrefixField) String() string {
	strSlice := []string{
		"name=" + `"` + s.AppName + `"`,
		"hostname=" + `"` + s.Hostname + `"`,
		"env=" + `"` + s.AppEnv + `"`,
		"branch=" + `"` + s.AppBranch + `"`,
		"version=" + `"` + s.AppVersion + `"`,
	}
	return strings.Join(strSlice, " ")
}

// Prefix 日志前缀
func (s *LoggerPrefixField) Prefix() []interface{} {
	ss := []string{
		"hostname=" + `"` + s.Hostname + `"`,
		"env=" + `"` + s.AppEnv + `"`,
		"branch=" + `"` + s.AppBranch + `"`,
		"version=" + `"` + s.AppVersion + `"`,
	}
	return []interface{}{
		"service", s.AppName,
		"app", strings.Join(ss, " "),
	}
}

// Config 配置
type Config interface {
	Watch(key string, o config.Observer) error

	// AppConfig APP配置
	AppConfig() *commonv1.App

	// ServerConfig 服务配置
	ServerConfig() *confv1.Server
	// HTTPConfig http配置
	HTTPConfig() *confv1.Server_HTTP
	// GRPCConfig grpc配置
	GRPCConfig() *confv1.Server_GRPC
	// BusinessAuthConfig APP验证配置
	BusinessAuthConfig() *confv1.Business_Auth
	// BaseSettingConfig APP设置配置
	BaseSettingConfig() *confv1.Base_Setting

	// Env app环境
	Env() commonv1.EnvEnum_Env
	// IsDebugMode 是否启用 调试模式
	IsDebugMode() bool
	// EnableLoggingConsole 是否启用 日志输出到控制台
	EnableLoggingConsole() bool
	// EnableLoggingFile 是否启用 日志输出到文件
	EnableLoggingFile() bool

	// LoggerConfigForConsole 日志配置 控制台
	LoggerConfigForConsole() *confv1.Log_Console
	// LoggerConfigForFile 日志配置 文件
	LoggerConfigForFile() *confv1.Log_File
	// DataConfig data配置
	DataConfig() *confv1.Data
	// MySQLConfig mysql配置
	MySQLConfig() *confv1.Data_MySQL
	// PostgresConfig postgres配置
	PostgresConfig() *confv1.Data_PSQL
	// RedisConfig redis配置
	RedisConfig() *confv1.Data_Redis
	// ConsulConfig consul配置
	ConsulConfig() *confv1.Base_Consul
	// JaegerTracerConfig jaeger tracer 配置
	JaegerTracerConfig() *confv1.Base_JaegerTracer
	// SnowflakeNodeIDConfig snowflake node
	SnowflakeNodeIDConfig() *configv1.NodeID
	// SnowflakeWorkerConfig snowflake worker 配置
	SnowflakeWorkerConfig() *confv1.Base_SnowflakeWorker

	// Close 关闭
	Close() error
}

// Engine 引擎模块、组件、单元
type Engine interface {
	// Config 配置
	Config

	// LoggerPrefixField 日志前缀字段
	LoggerPrefixField() *LoggerPrefixField
	// LoggerFileWriter 文件日志写手柄
	LoggerFileWriter() (io.Writer, error)
	// Logger 日志处理实例 runtime.caller.skip + 1
	// 用于 log.Helper 输出；例子：log.Helper.Info
	Logger() (log.Logger, []io.Closer, error)
	// LoggerHelper 日志处理实例 runtime.caller.skip + 2
	// 用于包含 log.Helper 输出；例子：func Info(){log.Helper.Info()}
	LoggerHelper() (log.Logger, []io.Closer, error)
	// LoggerMiddleware 日志处理实例 runtime.caller.skip - 1
	// 用于包含 http.Middleware(logging.Server)
	LoggerMiddleware() (log.Logger, []io.Closer, error)

	// GetMySQLGormDB mysql gorm 数据库
	GetMySQLGormDB() (*gorm.DB, error)
	// GetPostgresGormDB postgres gorm 数据库
	GetPostgresGormDB() (*gorm.DB, error)
	// GetRedisClient redis 客户端
	GetRedisClient() (*redis.Client, error)

	// SetRegistryType 设置 服务注册类型
	SetRegistryType(rt registrypkg.RegistryType)
	// RegistryType 服务注册类型
	RegistryType() registrypkg.RegistryType
	// GetConsulClient consul 客户端
	GetConsulClient() (*api.Client, error)

	// GetJaegerTraceExporter jaegerTrace
	GetJaegerTraceExporter() (*jaeger.Exporter, error)

	// GetAuthTokenRepo 验证Token工具
	GetAuthTokenRepo(redisCC *redis.Client) tokenutil.AuthTokenRepo

	// Close 关闭
	Close() error
}
