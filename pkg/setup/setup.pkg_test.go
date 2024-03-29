package setuppkg

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	logutil "github.com/ikaiguang/go-srv-kit/log"

	debugutil "github.com/ikaiguang/go-srv-kit/debug"
)

// go test -v ./pkg/setup/ -count=1 -test.run=TestSetup -conf=./../../testdata/configs
// 或使用 setup_main.pkg_test.go TestMain 配置的 testdata/configs
// go test -v ./pkg/setup/ -count=1 -test.run=TestSetup
func TestSetup(t *testing.T) {
	//engineHandler, err := New(WithConfigPath(confPath))
	engineHandler, err := New()
	if err != nil {
		t.Errorf("%+v\n", err)
		t.FailNow()
	}
	defer func() { _ = engineHandler.Close() }()

	ctx := context.Background()

	// env
	logutil.Infof("env = %v", engineHandler.Env())

	// debug
	debugutil.Println("*** | ==> debug util print")

	// 日志
	logutil.Info("*** | ==> log helper info")

	// db
	db, err := engineHandler.GetMySQLGormDB()
	require.Nil(t, err)
	require.NotNil(t, db)
	type DBRes struct {
		DBName string `gorm:"column:db_name"`
	}
	var dbRes DBRes
	err = db.WithContext(ctx).Raw("SELECT DATABASE() AS db_name").Scan(&dbRes).Error
	require.Nil(t, err)
	t.Logf("db res : %+v\n", dbRes)

	// redis
	redisCC, err := engineHandler.GetRedisClient()
	require.Nil(t, err)
	require.NotNil(t, redisCC)
	redisKey := "test-foo"
	redisValue := "test-bar"
	err = redisCC.Set(ctx, redisKey, redisValue, 0).Err()
	require.Nil(t, err)
	redisGotValue, err := redisCC.Get(ctx, redisKey).Result()
	require.Nil(t, err)
	require.Equal(t, redisValue, redisGotValue)
	t.Logf("redis res : %+v\n", redisGotValue)
}
