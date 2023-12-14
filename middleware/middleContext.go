package middleware

import (
	"github.com/casbin/casbin/v2"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type Config struct {
	SvcName string
	Redis   *redis.Redis
	Rbac    *casbin.Enforcer
}

type MiddleContext struct {
	// note: 中间件上下文
	//Dao   *dao.Dao
	SvcName string
	Redis   *redis.Redis
	Rbac    *casbin.Enforcer
}

func NewMiddlewareContext(in Config) *MiddleContext {
	return &MiddleContext{
		//Dao:   svc.Dao,
		SvcName: in.SvcName,
		Redis:   in.Redis,
		Rbac:    in.Rbac,
	}
}
