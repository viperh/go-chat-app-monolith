package routes

import (
	"go-chat-app-monolith/internal/api/controllers"

	"github.com/gin-gonic/gin"
)

func SetRoutes(g *gin.Engine, c *controllers.Controller) {

	authGroup := g.Group("/auth")

	authGroup.POST("/login", c.Login)
	authGroup.POST("/register", c.Register)

	wsGroup := g.Group("/ws")
	wsGroup.Use(c.Middleware.AuthRequired())

	authenticated := g.Group("/api")
	authenticated.Use(c.Middleware.AuthRequired())

	// TODO: add routes for crud operations users

	// TODO: add routes for crud operations rooms
}
