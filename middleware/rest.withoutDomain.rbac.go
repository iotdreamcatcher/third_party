package middleware

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"net/http"
	"third_party/commKey"
	"third_party/response"
)

type RestRbacInterceptorMiddleware struct {
	SvcName string
	Rbac    *casbin.Enforcer
}

func NewRestRbacInterceptorMiddleware(name string, rdb *redis.Redis, rbac *casbin.Enforcer) *RestRbacInterceptorMiddleware {
	return &RestRbacInterceptorMiddleware{
		SvcName: name,
		Rbac:    rbac,
	}
}

func (m *RestRbacInterceptorMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//租户权限个人验证、子账号验证个人与域内组权限
		// subect, object, action
		subect := r.Context().Value("UserId").(uint)
		object := r.RequestURI
		action := r.Method

		ok, err := m.checkPermission(fmt.Sprintf(commKey.RBAC_SUB, subect), object, action)
		if err != nil {
			logx.Errorf("checkPermission error: %v", err)
			CommonErrResponse(w, r, response.SERVER_WRONG)
			return
		}

		if !ok {
			CommonErrResponse(w, r, response.ACCESS_DENY)
			return
		}

		next(w, r)
	}
}

func (m *RestRbacInterceptorMiddleware) checkPermission(sub, obj, act string) (bool, error) {
	ok, err := m.Rbac.Enforce(sub, obj, act)
	if err != nil {
		return false, err
	}
	if !ok {
		return false, nil
	}
	return true, nil
}
