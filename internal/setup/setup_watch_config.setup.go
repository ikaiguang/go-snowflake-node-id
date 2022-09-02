package setup

import (
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/log"
	stdlog "log"
)

// watchConfigApp 监听配置 app
func (s *engines) watchConfigApp() (err error) {
	stdlog.Println("|*** 加载：监听配置：App")
	var observer = func(k string, v config.Value) {
		_ = s.logger.Log(log.LevelInfo,
			"watchConfigApp",
			"监听配置：数据有改动：key = "+k,
		)

		// app
		appConfig := s.AppConfig()
		if err := v.Scan(appConfig); err != nil {
			_ = s.logger.Log(log.LevelError,
				"watchConfigApp",
				"config.Value.Scan(appConfig) err : "+err.Error(),
			)
		}
	}
	if err = s.Watch("app", observer); err != nil {
		return err
	}
	return
}

// watchConfigData 监听配置 data
func (s *engines) watchConfigData() (err error) {
	stdlog.Println("|*** 加载：监听配置：Data")
	var observer = func(k string, v config.Value) {
		_ = s.logger.Log(log.LevelInfo,
			"watchConfigData",
			"监听配置：数据有改动：key = "+k,
		)

		// app
		dataConfig := s.DataConfig()
		if err := v.Scan(dataConfig); err != nil {
			_ = s.logger.Log(log.LevelError,
				"watchConfigData",
				"config.Value.Scan(appConfig) err : "+err.Error(),
			)
		}

		// 重新加载 mysql
		//if err := s.reloadMysqlGormDB(); err != nil {
		//	_ = s.logger.Log(log.LevelError,
		//		"watchConfigData",
		//		"reloadMysqlGormDB failed : "+err.Error(),
		//	)
		//}

		// 重新加载 postgres
		//if err := s.reloadPostgresGormDB(); err != nil {
		//	_ = s.logger.Log(log.LevelError,
		//		"watchConfigData",
		//		"reloadPostgresGormDB failed : "+err.Error(),
		//	)
		//}

		// 重新加载 redis
		//if err := s.reloadRedisClient(); err != nil {
		//	_ = s.logger.Log(log.LevelError,
		//		"watchConfigData",
		//		"reloadRedisClient failed : "+err.Error(),
		//	)
		//}
	}
	if err = s.Watch("data", observer); err != nil {
		return err
	}
	return
}
