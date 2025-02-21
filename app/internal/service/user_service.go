package service

import (
	"context"

	"github.com/deigo96/bpkp/app/config"
	"github.com/deigo96/bpkp/app/internal/entity"
	"github.com/deigo96/bpkp/app/internal/repository"
	"gorm.io/gorm"
)

type UserService interface {
	GetUsers(context.Context) ([]entity.UserResponse, error)
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

func (s *userService) GetUsers(c context.Context) ([]entity.UserResponse, error) {
	users, err := s.UserRepository.GetUsers(c, s.db)
	if err != nil {
		return nil, err
	}

	usersResponse := make([]entity.UserResponse, 0)
	if users == nil {
		return usersResponse, nil
	}

	for _, user := range users {
		userResponse := user.ToEntity()
		usersResponse = append(usersResponse, *userResponse)
	}

	return usersResponse, nil
}
