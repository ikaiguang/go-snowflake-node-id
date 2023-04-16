package snowflakeroute

import (
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	snowflakeservicev1 "github.com/ikaiguang/go-snowflake-node-id/api/snowflake-service/v1/services"
	snowflakeutil "github.com/ikaiguang/go-snowflake-node-id/business-util/snowflake"
	assemblers "github.com/ikaiguang/go-snowflake-node-id/internal/application/snowflake/assembler"
	services "github.com/ikaiguang/go-snowflake-node-id/internal/application/snowflake/service"
	srvs "github.com/ikaiguang/go-snowflake-node-id/internal/domain/snowflake/service"
	datas "github.com/ikaiguang/go-snowflake-node-id/internal/infra/snowflake/data"
	logutil "github.com/ikaiguang/go-srv-kit/log"
	"github.com/patrickmn/go-cache"
	stdlog "log"
	"time"

	"github.com/ikaiguang/go-snowflake-node-id/internal/setup"
)

// RegisterRoutes 注册路由
func RegisterRoutes(engineHandler setup.Engine, hs *http.Server, gs *grpc.Server) {
	stdlog.Println("|*** 注册路由：NodeID")

	// 数据库
	dbConn, err := engineHandler.GetMySQLGormDB()
	if err != nil {
		logutil.Fatal(err)
		return
	}

	// config
	nodeIDConfig := engineHandler.SnowflakeNodeIDConfig()

	// cache
	var locker snowflakeutil.Locker
	cacheHandler := cache.New(5*time.Minute, 10*time.Minute)
	locker = snowflakeutil.NewLockerFromCache(cacheHandler)
	if nodeIDConfig.EnableRedisLocker {
		redisCC, err := engineHandler.GetRedisClient()
		if err != nil {
			logutil.Fatal(err)
			return
		}
		locker = snowflakeutil.NewLockerFromRedis(redisCC)
	}

	// repos
	snowflakeWorkerRepo := datas.NewSnowflakeWorkerRepo(dbConn)

	snowflakeOptions := srvs.Options{
		MaxNodeID:    nodeIDConfig.MaxNodeId,
		IdleDuration: nodeIDConfig.IdleDuration.AsDuration(),
	}
	snowflakeSrv := srvs.NewSnowflakeSrv(snowflakeOptions, locker, snowflakeWorkerRepo)

	// 服务
	assembler := assemblers.NewAssembler()
	srv := services.NewWorker(assembler, snowflakeSrv)

	snowflakeservicev1.RegisterSrvSnowflakeWorkerHTTPServer(hs, srv)
	snowflakeservicev1.RegisterSrvSnowflakeWorkerServer(gs, srv)
}
