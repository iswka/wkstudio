package middleware

import (
	"net/http"
)

type AuthMiddleware struct {
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// JWT 中间件已经由 go-zero 框架自动处理
		// 这里可以添加额外的认证逻辑，比如检查用户状态等

		// 调用下一个处理器
		next(w, r)
	}
}
