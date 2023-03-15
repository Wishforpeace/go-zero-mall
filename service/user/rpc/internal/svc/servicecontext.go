package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"mall/service/user/model"
	"mall/service/user/rpc/internal/config"
)

// 注册服务上下文user model的依赖
type ServiceContext struct {
	Config config.Config

	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(conn, c.CacheRedis),
	}
}
