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

func (u *Employee) ToModel(totalPending int, leave []*LeaveRequest) *model.EmployeeResponse {
	var totalLeave, processedLeave int32
	for _, l := range leave {
		if l.Status == constant.PENDING {
			totalLeave += int32(l.TotalDays)
		}

		if l.Status == constant.APPROVED {
			processedLeave += int32(l.TotalDays)
		}
	}

	totalLeave = totalLeave + processedLeave + u.LeaveBalance

	return &model.EmployeeResponse{
		ID:                  u.ID,
		Name:                u.Name,
		NIP:                 u.Nip,
		Role:                u.RoleId.String(),
		LeaveBalance:        u.LeaveBalance,
		TotalPendingRequest: totalPending,
		Position:            u.Position,
		ProcessedLeave:      processedLeave,
		TotalLeave:          totalLeave,
		Department:          u.Department,
		CreatedAt:           u.CreatedAt.Format("2006-01-02 15:04:05"),
		CreatedBy:           u.CreatedBy,
		UpdatedAt:           u.UpdatedAt.Format("2006-01-02 15:04:05"),
		UpdatedBy:           u.UpdatedBy,
	}
}

func (e *Employee) ToEntity(req *model.CreateEmployeeRequest) {
	e.Nip = req.NIP
	e.Name = req.Name
	e.Password = req.Password
	e.RoleId = constant.GetRole(req.Role)
	e.LeaveBalance = req.LeaveBalance
	e.IsPns = req.IsPNS
	e.Position = req.Position
	e.Department = req.Department
}
