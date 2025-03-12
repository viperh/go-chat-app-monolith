package token

import (
	"github.com/golang-jwt/jwt/v5"
	"go-chat-app-monolith/internal/config"
	"time"
)

type Service struct {
	Secret string
}

func NewService(cfg *config.Config) *Service {
	return &Service{
		Secret: cfg.JwtSecret,
	}
}

func (s *Service) GenerateToken(userId uint) (string, error) {

	claims := jwt.MapClaims{
		"userId":    userId,
		"ExpiresAt": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(s.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil

}

func (s *Service) ValidateToken(tokenString string) (uint, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.Secret), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, err
	}

	userId := claims["userId"].(uint)

	return userId, nil
}
