package sockets

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

type Service struct {
	Upgrader *websocket.Upgrader
}

func NewService() *Service {
	upgrader := &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	return &Service{
		Upgrader: upgrader,
	}
}

func (s *Service) Upgrade(w gin.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	conn, err := s.Upgrader.Upgrade(*w, r, nil)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
