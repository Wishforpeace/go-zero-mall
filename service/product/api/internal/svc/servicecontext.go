package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"mall/service/product/api/internal/config"
	"mall/service/product/rpc/productclient"
)

type ServiceContext struct {
	Config     config.Config
	ProductRpc productclient.Product
}

// 注册服务上下文 product rpc 的依赖
func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		ProductRpc: productclient.NewProduct(zrpc.MustNewClient(c.ProductRpc)),
	}
}
