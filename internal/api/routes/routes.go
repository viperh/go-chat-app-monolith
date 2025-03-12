package routes

import "github.com/gin-gonic/gin"

func SetRoutes(g *gin.Engine, ) {
	g.POST("/login")
	g.POST("/register")
}
