package entity

import (
	"time"

	constant "github.com/deigo96/itineris/app/internal/const"
	"github.com/deigo96/itineris/app/internal/model"
)

type Employee struct {
	ID           int `gorm:"primaryKey"`
	Nip          string
	Name         string
	Password     string
	RoleId       constant.Role
	LeaveBalance int32
	IsPns        bool
	Position     string
	Department   string
	CreatedAt    *time.Time
	CreatedBy    string
	UpdatedAt    *time.Time
	UpdatedBy    string
}

func (u *Employee) ToModel(totalPending int) *model.EmployeeResponse {
	return &model.EmployeeResponse{
		ID:                  u.ID,
		Name:                u.Name,
		NIP:                 u.Nip,
		Role:                u.RoleId.String(),
		LeaveBalance:        u.LeaveBalance,
		TotalPendingRequest: totalPending,
		Position:            u.Position,
		Department:          u.Department,
		CreatedAt:           u.CreatedAt.Format("2006-01-02 15:04:05"),
		CreatedBy:           u.CreatedBy,
		UpdatedAt:           u.UpdatedAt.Format("2006-01-02 15:04:05"),
		UpdatedBy:           u.UpdatedBy,
	}
}

// func (u *Employee) CreateUserToEntity(c context.Context, user *model.CreateUserRequest) *Employee {
// 	return &Employee{
// 		Email:     user.Email,
// 		Password:  user.Password,
// 		Role:      constant.GetRole(user.Role),
// 		IsActive:  true,
// 		CreatedBy: "",
// 		UpdatedBy: "",
// 	}
// }
