package setuppkg

import (
	"testing"

	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/stretchr/testify/require"
)

// go test -v ./pkg/setup/ -count=1 -test.run=TestNewUpPackages
func TestNewUpPackages(t *testing.T) {
	// config
	var opts []config.Option
	opts = append(opts, config.WithSource(
		file.NewSource(confPath),
	))
	configHandler, err := NewConfiguration(opts...)
	if err != nil {
		t.Errorf("%+v\n", err)
		t.FailNow()
	}
	t.Log("*** | env：", configHandler.Env())

	// up
	upHandler := initEngine(configHandler)

	// db
	db, err := upHandler.GetMySQLGormDB()
	require.Nil(t, err)
	require.NotNil(t, db)

	// redis
	redisCC, err := upHandler.GetRedisClient()
	require.Nil(t, err)
	require.NotNil(t, redisCC)
}
