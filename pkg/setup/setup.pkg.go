package setuppkg

import (
	"flag"
	registrypkg "github.com/ikaiguang/go-snowflake-node-id/pkg/registry"
	stdlog "log"
)

// New 启动与配置
func New(opts ...Option) (engineHandler Engine, err error) {
	// parses the command-line flags
	flag.Parse()

	// 配置方式
	var (
		configHandler Config
	)
	switch {
	case _configConsulPath != "":
		// 初始化配置手柄
		configHandler, _, err = newConfigWithConsul(opts...)
		if err != nil {
			return engineHandler, err
		}
	default:
		// 初始化配置手柄
		configHandler, err = newConfigWithFiles(opts...)
		if err != nil {
			return engineHandler, err
		}
	}

	// 开始配置
	stdlog.Println("|==================== 配置程序 开始 ====================|")
	defer stdlog.Println("|==================== 配置程序 结束 ====================|")

	return newEngine(configHandler)
}

// newEngine 启动与配置
func newEngine(configHandler Config) (engineHandler Engine, err error) {
	// 初始化手柄
	setupHandler := initEngine(configHandler)

	// 设置调试工具
	if err = setupHandler.loadingDebugUtil(); err != nil {
		return engineHandler, err
	}

	// 设置日志工具
	if _, err = setupHandler.loadingLogHelper(); err != nil {
		return engineHandler, err
	}

	// mysql gorm 数据库
	if cfg := setupHandler.Config.MySQLConfig(); cfg != nil && cfg.Enable {
		if _, err = setupHandler.GetMySQLGormDB(); err != nil {
			return engineHandler, err
		}
	}

	// postgres gorm 数据库
	if cfg := setupHandler.Config.PostgresConfig(); cfg != nil && cfg.Enable {
		if _, err = setupHandler.GetPostgresGormDB(); err != nil {
			return engineHandler, err
		}
	}

	// redis 客户端
	if cfg := setupHandler.Config.RedisConfig(); cfg != nil && cfg.Enable {
		redisCC, err := setupHandler.GetRedisClient()
		if err != nil {
			return engineHandler, err
		}
		// 验证Token工具
		_ = setupHandler.GetAuthTokenRepo(redisCC)
	}

	// 服务注册
	setupHandler.SetRegistryType(registrypkg.RegistryTypeLocal)

	// consul 客户端
	if cfg := setupHandler.Config.ConsulConfig(); cfg != nil && cfg.Enable {
		_, err = setupHandler.GetConsulClient()
		if err != nil {
			return engineHandler, err
		}
	}

	// jaeger trace
	if cfg := setupHandler.Config.JaegerTracerConfig(); cfg != nil && cfg.Enable {
		_, err = setupHandler.GetJaegerTraceExporter()
		if err != nil {
			return engineHandler, err
		}
	}

	// 雪花算法
	//if cfg := setupHandler.Config.BaseSettingConfig(); cfg != nil && cfg.EnableSnowflakeWorker {
	//	err = setupHandler.loadingSnowflakeWorker()
	//	if err != nil {
	//		return engineHandler, err
	//	}
	//}

	// 监听配置 app
	//if err = setupHandler.watchConfigApp(); err != nil {
	//	return engineHandler, err
	//}

	// 监听配置 data
	//if err = setupHandler.watchConfigData(); err != nil {
	//	return engineHandler, err
	//}

	return setupHandler, err
}
