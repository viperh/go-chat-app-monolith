package users

import (
	"go-chat-app-monolith/internal/models"
	"go-chat-app-monolith/internal/pkg/provider"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	Provider *provider.Provider
}

func NewService(provider *provider.Provider) *Service {
	return &Service{Provider: provider}
}

func (s *Service) AddUser(user *models.User) error {
	_, err := s.GetUserByEmail(user.Email)
	if err == nil {
		return ErrUserExist
	}

	result := s.Provider.Db.Create(user)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (s *Service) DeleteUserById(id uint) error {
	res := s.Provider.Db.Delete(&models.User{}, id)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (s *Service) GetUserById(id uint) (*models.User, error) {
	user := &models.User{}
	res := s.Provider.Db.First(user, id)
	if res.Error != nil {
		return nil, res.Error
	}

	return user, nil
}

func (s *Service) GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	res := s.Provider.Db.Where("email = ?", email).First(user)
	if res.Error != nil {
		return nil, res.Error
	}

	return user, nil

}

func (s *Service) GetUserByUsername(username string) (*models.User, error) {
	user := &models.User{}
	res := s.Provider.Db.Where("username = ?", username).First(user)
	if res.Error != nil {
		return nil, res.Error
	}

	return user, nil

}

func (s *Service) UpdateUser(user *models.User) error {
	user, err := s.GetUserById(user.ID)
	if err != nil {
		return err
	}

	res := s.Provider.Db.Save(user)
	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (s *Service) CheckPassword(psswd string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(psswd))
	return err == nil
}
