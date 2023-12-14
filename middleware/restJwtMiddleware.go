/*
*

	@author: taco
	@Date: 2023/9/27
	@Time: 10:16

*
*/
package middleware

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"net/http"
	"third_party/cache_key"
	"third_party/commKey"
	utils "third_party/cryptography"
)

type JwtVerifyMiddleware struct {
	SvcName string
	Redis   *redis.Redis
	Rbac    *casbin.Enforcer
}

func NewJwtVerifyMiddleware(name string, rdb *redis.Redis, rbac *casbin.Enforcer) *JwtVerifyMiddleware {
	return &JwtVerifyMiddleware{
		SvcName: name,
		Redis:   rdb,
		Rbac:    rbac,
	}
}

type MiddleWareResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (m *JwtVerifyMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := new(MiddleWareResp)
		// TODO generate middleware implement function, delete after code implementation
		authToken := r.Header.Get(commKey.HANDER_AUTHORIZATION)
		if authToken == "" || len(authToken) <= 7 {
			w.WriteHeader(http.StatusOK)
			resp.Code = http.StatusUnauthorized
			resp.Msg = "token is empty"
			body, _ := json.Marshal(&resp)
			w.Write(body)
			return
		}
		token := authToken[7:]
		key := fmt.Sprintf(cache_key.ACCESS_TOKEN_KEY, m.SvcName, r.Header.Get(commKey.HANDER_ACCESSKEY))
		pubKey, err := m.Redis.Get(key)
		if err != nil {
			resp.Code = http.StatusUnauthorized
			resp.Msg = "accessKey is empty"
			body, _ := json.Marshal(&resp)
			w.Write(body)
			return
		}
		claims, err := utils.ParseToken(token, pubKey)
		if err != nil {
			resp.Code = http.StatusUnauthorized
			resp.Msg = "token is invalid"
			body, _ := json.Marshal(&resp)
			w.Write(body)
			return
		}
		//租户权限个人验证、子账号验证个人与域内组权限
		if claims.UserId != 1 {
			domain := claims.TenantID
			ok, err := m.checkPermission(fmt.Sprintf(commKey.RBAC_SUB, claims.UserId), fmt.Sprintf(commKey.RBAC_DOMAIN, claims.TenantID), r.RequestURI, r.Method)
			if err != nil {
				resp.Code = http.StatusInternalServerError
				resp.Msg = "checkPermission is invalid"
				body, _ := json.Marshal(&resp)
				w.Write(body)
				return
			}

			//租户验证组权限
			if !ok {
				if claims.UserId == claims.TenantID {
					//域设置为root的域
					domain = 1
				} else {
					resp.Code = http.StatusUnauthorized
					resp.Msg = "Permission verification failed"
					body, _ := json.Marshal(&resp)
					w.Write(body)
					return
				}
				//在验证组权限
				ok, err = m.checkPermission(fmt.Sprintf(commKey.RBAC_SUB, claims.UserId), fmt.Sprintf(commKey.RBAC_DOMAIN, domain), r.RequestURI, r.Method)
				if err != nil {
					resp.Code = http.StatusInternalServerError
					resp.Msg = "checkPermission is invalid"
					body, _ := json.Marshal(&resp)
					w.Write(body)
					return
				}
				if !ok {
					resp.Code = http.StatusUnauthorized
					resp.Msg = "Permission verification failed"
					body, _ := json.Marshal(&resp)
					w.Write(body)
					return
				}
			}
		}

		ctx := context.WithValue(r.Context(), commKey.CONTEXT_KEY_UID, claims.UserId)
		ctx = context.WithValue(ctx, commKey.CONTEXT_KEY_TENANTID, claims.TenantID)
		r = r.WithContext(ctx)
		// Passthrough to next handler if need
		next(w, r)
	}
}

func (m *JwtVerifyMiddleware) checkPermission(sub, domain, obj, act string) (bool, error) {
	ok, err := m.Rbac.Enforce(sub, domain, obj, act)
	if err != nil {
		return false, err
	}
	if !ok {
		return false, nil
	}
	return true, nil
}
