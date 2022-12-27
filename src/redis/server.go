package main

// 操作服务器
func opServer() {
	// 清空当前数据库中所有的 key
	redisCli.FlushDB(ctx)
	// 清空所有数据库里所有的 key
	redisCli.FlushAll(ctx)
	// 当前数据库的 key 数量
	redisCli.DBSize(ctx)
	// 当前服务器的时间
	redisCli.Time(ctx)
	// 获取连接到服务器的客户端连接列表
	redisCli.ClientList(ctx)
	// 关闭客户端连接
	redisCli.ClientKill(ctx, "111")
	// 获取连接的名称
	redisCli.ClientGetName(ctx)
	// 获取服务器配置
	redisCli.ConfigGet(ctx, "slowlog-max-len")
	// 修改服务器的配置，不需要重启
	redisCli.ConfigSet(ctx, "slowlog-max-len", "10086")
}
