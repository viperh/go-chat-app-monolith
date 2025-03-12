package controllers

import (
	"github.com/gin-gonic/gin"
	"go-chat-app-monolith/internal/api/controllers/dto"
	"go-chat-app-monolith/internal/models"
	"go-chat-app-monolith/internal/pkg/token"
	"go-chat-app-monolith/internal/pkg/users"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type Controller struct {
	UserService  *users.Service
	TokenService *token.Service
}

func NewController(us *users.Service, tokenService *token.Service) *Controller {
	return &Controller{
		UserService:  us,
		TokenService: tokenService,
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

	token, err := c.TokenService.GenerateToken(user.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resObj := gin.H{
		"token":  token,
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
