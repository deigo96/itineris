package model

import (
	"time"

	constant "github.com/deigo96/bpkp/app/internal/const"
	"github.com/deigo96/bpkp/app/internal/entity"
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

func (u *Users) ToEntity() *entity.UserResponse {
	return &entity.UserResponse{
		Email:     u.Email,
		IsActive:  u.IsActive,
		Role:      u.Role.String(),
		CreatedAt: u.CreatedAt.String(),
		CreatedBy: u.CreatedBy,
		UpdatedAt: u.UpdatedAt.String(),
		UpdatedBy: u.UpdatedBy,
	}
}
