package entity

import (
	"context"
	"time"

	constant "github.com/deigo96/itineris/app/internal/const"
	"github.com/deigo96/itineris/app/internal/model"
)

type Employee struct {
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

func (u *Employee) ToModel() *model.EmployeeResponse {
	return &model.EmployeeResponse{
		Email:     u.Email,
		IsActive:  u.IsActive,
		Role:      u.Role.String(),
		CreatedAt: u.CreatedAt.String(),
		CreatedBy: u.CreatedBy,
		UpdatedAt: u.UpdatedAt.String(),
		UpdatedBy: u.UpdatedBy,
	}
}

func (u *Employee) CreateUserToEntity(c context.Context, user *model.CreateUserRequest) *Employee {
	return &Employee{
		Email:     user.Email,
		Password:  user.Password,
		Role:      constant.GetRole(user.Role),
		IsActive:  true,
		CreatedBy: "",
		UpdatedBy: "",
	}
}
