package sockets

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Service struct {
	mu       sync.Mutex
	Upgrader *websocket.Upgrader
	Conns    map[uint]*websocket.Conn
}

func NewService() *Service {
	upgrader := &websocket.Upgrader{}
	return &Service{
		Upgrader: upgrader,
	}
}

func (s *Service) Upgrade(w gin.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	conn, err := s.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		return nil, err
	}

	return conn, nil

}

func (s *Service) AddConn(userId uint, conn *websocket.Conn) {
	s.mu.Lock()
	s.Conns[userId] = conn
	s.mu.Unlock()
}

func (s *Service) RemoveConn(userId uint) {
	s.mu.Lock()
	delete(s.Conns, userId)
	s.mu.Unlock()
}

func (s *Service) HandleConn(userId uint, conn *websocket.Conn) {
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			continue
		}

		for _, c := range s.Conns {
			c.WriteMessage(websocket.TextMessage, msg)
		}
	}
}

func (s *Service) Broadcast(origin *websocket.Conn, msg []byte) {
	for _, conn := range s.Conns {
		err := conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			origin.WriteMessage(websocket.TextMessage, []byte("Error broadcasting message: "+err.Error()))
		}
	}
}
