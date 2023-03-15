package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

// 添加mysql服务配置，CacheRedis服务配置的实力化
type Config struct {
	zrpc.RpcServerConf
	Mysql struct {
		DataSource string
	}

	CacheRedis cache.CacheConf

	Salt string
}
