package service

import (
	"InjectGo-Workshop/api/repository"
	"InjectGo-Workshop/model"
)

type UserService interface {
	CreateUser(user *model.User) (*model.User, error)
	GetUsers() (*[]model.User, error)
	GetUserByID(userID uint) (*model.User, error)
}

type IUserService struct {
	UserRepository repository.UserRepository
}

func NewUserService(UserRepository repository.UserRepository) UserService {
	return &IUserService{
		UserRepository: UserRepository,
	}
}

func (s *IUserService) CreateUser(user *model.User) (*model.User, error) {
	return s.UserRepository.CreateUser(user)
}

func (s *IUserService) GetUsers() (*[]model.User, error) {
	return s.UserRepository.GetUsers()
}

func (s *IUserService) GetUserByID(userID uint) (*model.User, error) {
	user, err := s.UserRepository.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}
