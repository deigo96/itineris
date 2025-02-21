package repository

import (
	"context"

	"github.com/deigo96/bpkp/app/internal/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUsers(context.Context, *gorm.DB) ([]model.Users, error)
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) GetUsers(c context.Context, db *gorm.DB) ([]model.Users, error) {
	users := []model.Users{}
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
