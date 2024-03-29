package setuppkg

import (
	"testing"

	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
)

// go test -v ./pkg/setup/ -count=1 -test.run=TestNewConfiguration
func TestNewConfiguration(t *testing.T) {
	var opts []config.Option
	opts = append(opts, config.WithSource(
		file.NewSource(confPath),
	))
	handler, err := NewConfiguration(opts...)
	if err != nil {
		t.Errorf("%+v\n", err)
		t.FailNow()
	}

	t.Log("*** | env：", handler.Env())
	t.Logf("*** | AppConfig：%+v\n", handler.AppConfig())
}
