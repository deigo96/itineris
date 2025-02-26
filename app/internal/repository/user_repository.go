package repository

import (
	"context"

	"github.com/deigo96/itineris/app/internal/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUsers(context.Context, *gorm.DB) ([]entity.Users, error)
	CreateUser(context.Context, *gorm.DB, *entity.Users) (*entity.Users, error)
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) GetUsers(c context.Context, db *gorm.DB) ([]entity.Users, error) {
	users := []entity.Users{}
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepository) CreateUser(c context.Context, db *gorm.DB, user *entity.Users) (*entity.Users, error) {
	if err := db.Create(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
