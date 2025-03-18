package sockets

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go-chat-app-monolith/internal/pkg/token"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type Service struct {
	mu         *sync.Mutex
	Upgrader   *websocket.Upgrader
	Conns      map[string]*websocket.Conn
	JwtService *token.Service
}

func NewService(jwtService *token.Service) *Service {
	upgrader := &websocket.Upgrader{}
	websocket.DefaultDialer.HandshakeTimeout = 3 * time.Second

	return &Service{
		Upgrader:   upgrader,
		mu:         &sync.Mutex{},
		Conns:      make(map[string]*websocket.Conn),
		JwtService: jwtService,
	}

}

func (s *Service) Upgrade(w gin.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	conn, err := s.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func (s *Service) AuthorizeClient(conn *websocket.Conn) {
	msg := &AuthMessage{}
	for {
		err := conn.ReadJSON(msg)
		if err != nil {
			err = conn.WriteMessage(websocket.TextMessage, []byte("You failed to authorize yourself!"))
			return
		}
		userId, err := s.JwtService.ValidateToken(msg.Token)

		_ = conn.WriteMessage(websocket.TextMessage, []byte(`You are now authorized as: `+strconv.Itoa(int(userId))))
		s.Conns[msg.Token] = conn
		return
	}
}

func (s *Service) HandleConn(conn *websocket.Conn) {
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		go s.Broadcast(conn, msg)
	}
	conn.Close()
}

func (s *Service) Broadcast(origin *websocket.Conn, msg []byte) {
	for _, conn := range s.Conns {
		if conn == origin {
			continue
		}
		err := conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			_ = origin.WriteMessage(websocket.TextMessage, []byte("Error broadcasting message: "+err.Error()))
		}
	}
}
