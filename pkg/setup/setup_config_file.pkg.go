package setuppkg

import (
	"flag"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	stdlog "log"
)

// newConfigWithFiles 初始化配置手柄
func newConfigWithFiles(setupOpts ...Option) (Config, error) {
	if !flag.Parsed() {
		flag.Parse()
	}

	// 启动选项
	setupOpt := &options{}
	for i := range setupOpts {
		setupOpts[i](setupOpt)
	}

	stdlog.Println("|==================== 加载配置文件 开始 ====================|")
	defer stdlog.Println()
	defer stdlog.Println("|==================== 加载配置文件 结束 ====================|")
	// 配置路径
	confPath := _configFilepath
	if setupOpt.configPath != "" {
		confPath = setupOpt.configPath
	} else if confPath == "" {
		confPath = _defaultConfigFilepath
	}

	var opts []config.Option
	stdlog.Println("|*** 加载：配置文件路径: ", confPath)
	opts = append(opts, config.WithSource(file.NewSource(confPath)))
	return NewConfiguration(opts...)
}
