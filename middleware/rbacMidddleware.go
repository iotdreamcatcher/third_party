package middleware

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logc"
	"google.golang.org/grpc"
	"strings"
	"third_party/response"
)

// note: 基于grpc的中间件，实现rbac鉴权认证

func (SvcCtx *MiddleContext) RbacUnaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

	result := new(Resp)
	result.Path = info.FullMethod
	result.RequestID = ctx.Value("RequestID").(string)

	UserId := ctx.Value("UserId")
	TenantID := ctx.Value("TenantID")
	objAndAct := strings.Split(info.FullMethod, "/")
	if len(objAndAct) > 2 {

		sub := fmt.Sprintf("UID_%d", UserId.(uint))
		domain := fmt.Sprintf("TID_%d", TenantID.(uint))
		obj := objAndAct[1]
		act := objAndAct[2]

		// note: rbac打印
		logc.Infof(ctx, "sub: %s, domain: %s, obj: %s, act: %s", sub, domain, obj, act)

		ok, err := SvcCtx.Rbac.Enforce(sub, domain, obj, act)
		if err != nil {
			result.Code = response.SERVER_WRONG
			result.Msg = response.StatusText(response.SERVER_WRONG)
			logc.Errorf(ctx, "权限验证出错: %v", err)
			return result, nil
		}
		if !ok {
			// note: 设置参数到标准返回上
			// note: 标准返回根据服务名获取对应的返回结构体，并映射到proto message
			result.Code = response.ACCESS_DENY
			result.Msg = response.StatusText(response.ACCESS_DENY)
			return result, nil
		}
	} else {
		result.Code = response.AUTHORIZATION_NOT_FOUND
		result.Msg = response.StatusText(response.AUTHORIZATION_NOT_FOUND)
		return result, nil
	}

	return handler(ctx, req)
}
func RbacStreamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	// TODO: fill your logic here
	return handler(srv, ss)
}
