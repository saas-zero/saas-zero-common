package middleware

import (
	"context"
	"net/http"
	"strings"

	ryjwt "github.com/saas-zero/saas-zero-common/pkg/jwt"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type JWTMiddleware struct {
	Conf *ryjwt.TokenConf
}

func NewJWTMiddleware(conf *ryjwt.TokenConf) *JWTMiddleware {
	return &JWTMiddleware{Conf: conf}
}

func (m *JWTMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := m.extractToken(r)
		if token == "" {
			httpx.Error(w, NewUnauthorizedError("未提供认证令牌"))
			return
		}

		userId, err := ryjwt.Valid(m.Conf, "userId", token)
		if err != nil {
			httpx.Error(w, NewUnauthorizedError("无效的认证令牌"))
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, UserIdKey, userId)
		next(w, r.WithContext(ctx))
	}
}

func (m *JWTMiddleware) extractToken(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	if authHeader != "" {
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) == 2 && strings.ToLower(parts[0]) == "bearer" {
			return parts[1]
		}
	}

	token := r.URL.Query().Get("token")
	if token != "" {
		return token
	}

	cookie, err := r.Cookie("token")
	if err == nil {
		return cookie.Value
	}

	return ""
}