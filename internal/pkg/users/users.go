package users

import "gorm.io/gorm"

type Service struct {
	Db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{Db: db}
}

func (s *Service) AddUser(CreateUserReq) error {
	return nil
}

func (s *Service) DeleteUserById(DeleteUserReq) error {
	return nil
}

func (s *Service) GetuserById(GetUserByIdReq) error {
	return nil
}

func (s *Service) GetUserByEmail(GetUserByEmailReq) error {
	return nil
}

func (s *Service) UpdateUser(UpdateUserReq) error {
	return nil
}

func (s *Service) Login(LoginReq) error {
	return nil
}


