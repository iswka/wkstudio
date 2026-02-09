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
		// JWT middleware is automatically handled by go-zero framework
		// Additional authentication logic can be added here, such as checking user status

		// Call the next handler
		next(w, r)
	}
}
