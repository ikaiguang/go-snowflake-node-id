package setuppkg

import (
	"github.com/ikaiguang/go-snowflake-node-id/testdata"
	"os"
	"testing"
)

// confPath 配置目录
const confPath = "./../../testdata/configs"

func TestMain(m *testing.M) {
	_configFilepath = testdata.ConfigPath()

	// 初始化必要逻辑
	os.Exit(m.Run())
}
