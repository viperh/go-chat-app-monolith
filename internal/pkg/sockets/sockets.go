package sockets

import (
	"go-chat-app-monolith/internal/config"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Service struct {
	Upgrader *websocket.Upgrader
}

func NewService(cfg *config.Config) *Service {
	return &Service{
		Upgrader: &websocket.Upgrader{}
	}
}

func (s *Service) Upgrade(ctx *gin.Context) (*websocket.Conn, error) {
	ws, err := s.Upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		return nil, err
	}
	return ws, nil
}
