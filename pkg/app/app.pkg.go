package apppkg

import (
	commonv1 "github.com/ikaiguang/go-snowflake-node-id/api/common/v1"
	"strings"
)

const (
	_appIDSep      = ":"
	_configPathSep = "/"
)

// IsDebugMode ...
func IsDebugMode(appEnv commonv1.EnvEnum_Env) bool {
	switch appEnv {
	case commonv1.EnvEnum_DEVELOP, commonv1.EnvEnum_TESTING:
		return true
	default:
		return false
	}
}

// ID 程序ID
// 例：go-srv-services/DEVELOP/main/v1.0.0/user-service
func ID(appConfig *commonv1.App) string {
	return appIdentifier(appConfig, _appIDSep)
}

// ConfigPath 配置路径；用于配置中心，如：consul、etcd、...
// @result = app.BelongTo + "/" + app.Env + "/" + app.Branch + "/" + app.Version + "/" + app.Name
// 例：go-srv-services/DEVELOP/main/v1.0.0/user-service
func ConfigPath(appConfig *commonv1.App) string {
	return appIdentifier(appConfig, _configPathSep)
}

// appIdentifier app 唯一标准
// @result = app.BelongTo + "/" + app.Env + "/" + app.Branch + "/" + app.Version + "/" + app.Name
// 例：go-srv-services/DEVELOP/main/v1.0.0/user-service
func appIdentifier(appConfig *commonv1.App, sep string) string {
	var ss = make([]string, 0, 5)
	if appConfig.BelongTo != "" {
		ss = append(ss, appConfig.BelongTo)
	}
	ss = append(ss, ParseEnv(appConfig.Env).String())
	if appConfig.EnvBranch != "" {
		branchString := strings.Replace(appConfig.EnvBranch, " ", ":", -1)
		ss = append(ss, branchString)
	}
	if appConfig.Version != "" {
		ss = append(ss, appConfig.Version)
	}
	if appConfig.Name != "" {
		ss = append(ss, appConfig.Name)
	}
	return strings.Join(ss, sep)
}

// ParseEnv ...
func ParseEnv(appEnv string) (envEnum commonv1.EnvEnum_Env) {
	envInt32, ok := commonv1.EnvEnum_Env_value[strings.ToUpper(appEnv)]
	if ok {
		envEnum = commonv1.EnvEnum_Env(envInt32)
	}
	if envEnum == commonv1.EnvEnum_UNKNOWN {
		envEnum = commonv1.EnvEnum_PRODUCTION
		return envEnum
	}
	return envEnum
}
