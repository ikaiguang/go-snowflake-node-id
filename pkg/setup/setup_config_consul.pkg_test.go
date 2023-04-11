package setuppkg

import (
	"testing"
)

// go test -v ./pkg/setup/ -count=1 -test.run=TestNew_newConfigWithConsul -conf-consul=./../../app/admin-service/configs/consul
func TestNew_newConfigWithConsul(t *testing.T) {
	var opts []Option

	// 在 初始化Consul配置中心 结束 前没有错误即为测试成功
	_, _, _ = newConfigWithConsul(opts...)
	//if err != nil {
	//t.Logf("%+v\n", err)
	//t.FailNow()
	//}

	//t.Log("*** | env：", handler.Env())
	//t.Logf("*** | AppConfig：%+v\n", handler.AppConfig())
}
