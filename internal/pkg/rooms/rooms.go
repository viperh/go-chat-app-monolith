package rooms

import (
	"go-chat-app-monolith/internal/models"
	"go-chat-app-monolith/internal/pkg/provider"
)

type Service struct {
	Provider *provider.Provider
}

func NewService(prov *provider.Provider) *Service {
	return &Service{
		Provider: prov,
	}
}

func (s *Service) CreateRoom() error {
	res := s.Provider.Db.Create(&models.Room{})
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (s *Service) DeleteRoomById(id uint) error {
	res := s.Provider.Db.Delete(&models.Room{}, id)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (s *Service) GetRoomById(id uint) (*models.Room, error) {
	room := &models.Room{}
	res := s.Provider.Db.First(room, id)
	if res.Error != nil {
		return nil, res.Error
	}
	return room, nil
}

func (s *Service) GetRoomsByUserId(id uint) ([]*models.Room, error) {
	var rooms []*models.Room
	res := s.Provider.Db.Find(&rooms).Where("userId = ?", id)
	if res.Error != nil {
		return nil, res.Error
	}

	return rooms, nil
}
