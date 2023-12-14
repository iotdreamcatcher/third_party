package middleware

import (
	"context"
	"net/http"
)

type PathHttpInterceptorMiddleware struct {
}

func NewPathHttpInterceptorMiddleware() *PathHttpInterceptorMiddleware {
	return &PathHttpInterceptorMiddleware{}
}

func (m *PathHttpInterceptorMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "FullMethod", r.URL.Path)
		ctx = context.WithValue(ctx, "RequestURI", r.RequestURI)
		r = r.WithContext(ctx)
		next(w, r)
	}
}
