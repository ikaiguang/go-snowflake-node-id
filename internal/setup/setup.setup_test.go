package setup

import (
	debugutil "github.com/ikaiguang/go-srv-kit/debug"
	logutil "github.com/ikaiguang/go-srv-kit/log"
	"github.com/stretchr/testify/require"
	"testing"
)

// go test -v ./internal/setup/ -count=1 -test.run=TestGetEngine -conf=./../../configs
func TestGetEngine(t *testing.T) {
	engineHandler, err := GetEngine()
	require.Nil(t, err)
	require.NotNil(t, engineHandler)
	engineHandler, err = GetEngine()
	require.Nil(t, err)
	require.NotNil(t, engineHandler)

	// env
	logutil.Infof("env = %v", engineHandler.Env())

	// debug
	debugutil.Println("*** | ==> debug util print")

	// 日志
	logutil.Info("*** | ==> log helper info")
}
