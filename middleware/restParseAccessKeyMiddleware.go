package middleware

import (
	"context"
	"net/http"
	"third_party/commKey"
)

type ParseAccessKeyMiddleware struct {
}

func NewParseAccessKeyMiddleware() *ParseAccessKeyMiddleware {
	return &ParseAccessKeyMiddleware{}
}

func (m *ParseAccessKeyMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation
		accessKey := r.Header.Get(commKey.HANDER_ACCESSKEY)
		ctx := context.WithValue(r.Context(), "accessKey", accessKey)
		r = r.WithContext(ctx)
		// Passthrough to next handler if need
		next(w, r)
	}
}
