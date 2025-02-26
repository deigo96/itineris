package entity

import (
	"context"
	"time"

	constant "github.com/deigo96/itineris/app/internal/const"
	"github.com/deigo96/itineris/app/internal/model"
)

type Users struct {
	ID        int `gorm:"primaryKey"`
	Email     string
	Password  string
	Role      constant.Role
	IsActive  bool
	CreatedAt *time.Time
	CreatedBy string
	UpdatedAt *time.Time
	UpdatedBy string
}

func (u *Users) ToModel() *model.UserResponse {
	return &model.UserResponse{
		Email:     u.Email,
		IsActive:  u.IsActive,
		Role:      u.Role.String(),
		CreatedAt: u.CreatedAt.String(),
		CreatedBy: u.CreatedBy,
		UpdatedAt: u.UpdatedAt.String(),
		UpdatedBy: u.UpdatedBy,
	}
}

func (u *Users) CreateUserToEntity(c context.Context, user *model.CreateUserRequest) *Users {
	return &Users{
		Email:     user.Email,
		Password:  user.Password,
		Role:      constant.GetRole(user.Role),
		IsActive:  true,
		CreatedBy: "",
		UpdatedBy: "",
	}
}
