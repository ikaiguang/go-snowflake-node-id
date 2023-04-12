package servers

// getAuthWhiteList 验证白名单
func getAuthWhiteList() map[string]struct{} {
	// 白名单
	whiteList := make(map[string]struct{})

	// 测试
	whiteList["/kit.api.pingservicev1.SrvPing/Ping"] = struct{}{}
	whiteList["/snowflake.service.api.snowflakeservicev1.SrvSnowflakeWorker/GetNodeId"] = struct{}{}
	whiteList["/snowflake.service.api.snowflakeservicev1.SrvSnowflakeWorker/ExtendNodeId"] = struct{}{}

	return whiteList
}
