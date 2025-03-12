package routes

import (
	"github.com/gin-gonic/gin"
	"go-chat-app-monolith/internal/api/controllers"
)

func SetRoutes(g *gin.Engine, c *controllers.Controller) {
	g.POST("/login", c.Login)
	g.POST("/register", c.Register)
}
