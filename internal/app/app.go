package app

import (
	"fmt"
	"go-chat-app-monolith/internal/api/controllers"
	"go-chat-app-monolith/internal/api/middlewares"
	"go-chat-app-monolith/internal/api/routes"
	"go-chat-app-monolith/internal/config"
	"go-chat-app-monolith/internal/pkg/provider"
	"go-chat-app-monolith/internal/pkg/token"
	"go-chat-app-monolith/internal/pkg/users"

	"github.com/gin-gonic/gin"
)

type App struct {
	Config       *config.Config
	AppProvider  *provider.Provider
	UserService  *users.Service
	TokenService *token.Service
	Controller   *controllers.Controller
	Engine       *gin.Engine
}

func NewApp() *App {
	cfg := config.MustLoad()
	prov := provider.NewProvider(cfg)
	userService := users.NewService(prov)
	jwtService := token.NewService(cfg)
	mw := middlewares.NewMiddleware(jwtService)
	controller := controllers.NewController(userService, jwtService, mw)
	engine := gin.Default()

	// test-only
	engine.Use(gin.Logger())

	routes.SetRoutes(engine, controller)

	return &App{
		Config:       cfg,
		AppProvider:  prov,
		UserService:  userService,
		TokenService: jwtService,
		Controller:   controller,
		Engine:       engine,
	}
}

func (a *App) Run() {
	gin.SetMode(gin.ReleaseMode)
	address := fmt.Sprintf(":%s", a.Config.ApiPort)
	err := a.Engine.Run(address)
	if err != nil {
		panic(err)
	}
}
