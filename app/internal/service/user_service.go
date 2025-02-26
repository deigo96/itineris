package service

import (
	"context"

	"github.com/deigo96/itineris/app/config"
	"github.com/deigo96/itineris/app/internal/model"
	"github.com/deigo96/itineris/app/internal/repository"
	"gorm.io/gorm"
)

type UserService interface {
	GetUsers(context.Context) ([]model.UserResponse, error)
}

type userService struct {
	UserRepository repository.UserRepository
	db             *gorm.DB
	config         *config.Config
}

func NewUserService(db *gorm.DB, config *config.Config) UserService {
	return &userService{
		UserRepository: repository.NewUserRepository(),
		db:             db,
		config:         config,
	}
}

func (s *userService) GetUsers(c context.Context) ([]model.UserResponse, error) {
	users, err := s.UserRepository.GetUsers(c, s.db)
	if err != nil {
		return nil, err
	}

	usersResponse := make([]model.UserResponse, 0)
	if users == nil {
		return usersResponse, nil
	}

	for _, user := range users {
		userResponse := user.ToModel()
		usersResponse = append(usersResponse, *userResponse)
	}

	return usersResponse, nil
}

// func (s *userService) CreateUser(c context.Context, user *model.CreateUserRequest) (*model.UserResponse, error) {

// }

// func (s *userService) constructCreateUser(c context.Context, user *model.CreateUserRequest) (*model.Users, error) {

// }
