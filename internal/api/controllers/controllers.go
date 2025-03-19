package controllers

import (
	"github.com/gin-gonic/gin"
	"go-chat-app-monolith/internal/api/controllers/dto"
	"go-chat-app-monolith/internal/api/middlewares"
	"go-chat-app-monolith/internal/models"
	"go-chat-app-monolith/internal/pkg/sockets"
	"go-chat-app-monolith/internal/pkg/token"
	"go-chat-app-monolith/internal/pkg/users"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type Controller struct {
	UserService   *users.Service
	TokenService  *token.Service
	Middleware    *middlewares.Middleware
	SocketGateway *sockets.Service
}

func NewController(us *users.Service, tokenService *token.Service, middleware *middlewares.Middleware, socketService *sockets.Service) *Controller {
	return &Controller{
		UserService:   us,
		TokenService:  tokenService,
		Middleware:    middleware,
		SocketGateway: socketService,
	}
}

func (c *Controller) Login(ctx *gin.Context) {
	req := &dto.LoginReq{}
	err := ctx.ShouldBindBodyWithJSON(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := c.UserService.GetUserByUsername(req.Username)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ret := c.UserService.CheckPassword(req.Password, user.Password)
	if !ret {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	tkn, err := c.TokenService.GenerateToken(user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resObj := gin.H{
		"token":  tkn,
		"userId": user.ID,
	}

	ctx.JSON(http.StatusOK, resObj)

}

func (c *Controller) Register(ctx *gin.Context) {
	req := &dto.RegisterReq{}

	err := ctx.ShouldBindBodyWithJSON(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user := &models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: string(hash),
	}

	err = c.UserService.AddUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User created successfully"})

}

func (c *Controller) GetUserById(ctx *gin.Context) {
	req := &dto.GetUserByIdReq{}

	err := ctx.ShouldBindBodyWithJSON(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := c.UserService.GetUserById(req.ID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	res := &dto.GetUserByIdRes{}
	res.ID = user.ID
	res.Username = user.Username
	res.Email = user.Email

	ctx.JSON(http.StatusOK, gin.H{"user": user})

}

func (c *Controller) UpdateUser(ctx *gin.Context) {
	req := &dto.UpdateUserReq{}

	err := ctx.ShouldBindBodyWithJSON(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := c.UserService.GetUserById(req.ID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	user.Username = req.Username
	user.Email = req.Email
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user.Password = string(hash)

	err = c.UserService.UpdateUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func (c *Controller) UpgradeToWs(ctx *gin.Context) {
	conn, err := c.SocketGateway.Upgrade(ctx.Writer, ctx.Request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.SocketGateway.AuthorizeClient(conn)
	go c.SocketGateway.HandleConn(conn)

}
