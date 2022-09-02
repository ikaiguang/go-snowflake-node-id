package services

import (
	"github.com/google/wire"

	pingsrv "github.com/ikaiguang/go-snowflake-node-id/internal/application/service/ping"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(
	// ping 服务
	pingsrv.NewPingService,
)
