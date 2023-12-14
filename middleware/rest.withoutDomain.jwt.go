package middleware

import (
	"context"
	"fmt"
	"github.com/iotdreamcatcher/third_party/cache_key"
	"github.com/iotdreamcatcher/third_party/commKey"
	"github.com/iotdreamcatcher/third_party/jwts"
	"github.com/iotdreamcatcher/third_party/response"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"net/http"
)

type RestJwtAuthInterceptorMiddleware struct {
	SvcName string
	Redis   *redis.Redis
}

func NewRestJwtAuthInterceptorMiddleware(name string, rdb *redis.Redis) *RestJwtAuthInterceptorMiddleware {
	return &RestJwtAuthInterceptorMiddleware{
		SvcName: name,
		Redis:   rdb,
	}
}

func (m *RestJwtAuthInterceptorMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authToken := r.Header.Get(commKey.HANDER_AUTHORIZATION)
		if authToken == "" || len(authToken) <= 7 {
			CommonErrResponse(w, r, response.AUTHORIZATION_NOT_FOUND)
			return
		}
		token := authToken[7:]
		key := fmt.Sprintf(cache_key.ACCESS_TOKEN_KEY, m.SvcName, r.Header.Get(commKey.HANDER_ACCESSKEY))
		logx.Infof("key: %v", key)
		pubKey, err := m.Redis.Get(key)
		if err != nil {
			logx.Errorf("pubKey error: %v", err)
			CommonErrResponse(w, r, response.ACCESSKEY_NOT_FOUND)
			return
		}
		if pubKey == "" || len(pubKey) <= 0 {
			logx.Infof("pubKey is empty，accessToken is expired")
			CommonErrResponse(w, r, response.ACCESS_EXPIRED)
			return
		}
		claims, err := jwts.JwtWithoutDomainParseToken(token, pubKey)
		if err != nil {
			logx.Errorf("ParseToken error: %v", err)
			CommonErrResponse(w, r, response.ACCESS_EXPIRED)
			return
		}
		ctx := context.WithValue(r.Context(), "UserId", claims.UserId)
		r = r.WithContext(ctx)
		next(w, r)
	}
}
