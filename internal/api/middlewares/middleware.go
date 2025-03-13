package middlewares

import (
	"go-chat-app-monolith/internal/pkg/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Middleware struct {
	JwtService *token.Service
}

func NewMiddleware(jwtService *token.Service) *Middleware {
	return &Middleware{
		JwtService: jwtService,
	}
}

func (m *Middleware) AuthRequired() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			ctx.Abort()
		}

		userId, err := m.JwtService.ValidateToken(tokenString)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			ctx.Abort()
		}

		ctx.Keys["userId"] = userId

		ctx.Next()

	}
}
