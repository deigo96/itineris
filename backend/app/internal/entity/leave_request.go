package entity

import (
	"time"

	constant "github.com/deigo96/itineris/app/internal/const"
	"github.com/deigo96/itineris/app/internal/model"
	"github.com/deigo96/itineris/app/internal/util"
)

type LeaveRequest struct {
	ID            int `gorm:"primaryKey"`
	EmployeeID    int
	Status        constant.Status
	LeaveType     int
	StartDate     time.Time
	EndDate       time.Time
	Reason        string
	RejectionNote string
	CreatedAt     *time.Time
	CreatedBy     string
	UpdatedAt     *time.Time
	UpdatedBy     string
	TotalDays     int
}

type UpdateLeaveRequest struct {
	Status        constant.Status
	RejectionNote string
	UpdatedBy     string
}

func (l *LeaveRequest) ToEntity(req *model.LeaveRequestRequest) {
	startDate, _ := util.ParseStringToTime(req.StartDate)
	endDate, _ := util.ParseStringToTime(req.EndDate)
	status := constant.Status(constant.PENDING)

	l.StartDate = startDate
	l.EndDate = endDate
	l.Reason = req.Reason
	l.Status = status
	l.LeaveType = req.LeaveType
}

func (l *LeaveRequest) ToModel(leaveType string) *model.LeaveRequestResponse {
	return &model.LeaveRequestResponse{
		ID:            l.ID,
		EmployeeID:    l.EmployeeID,
		Status:        l.Status.String(),
		StartDate:     l.StartDate.Format("2006-01-02"),
		EndDate:       l.EndDate.Format("2006-01-02"),
		Reason:        l.Reason,
		RejectionNote: l.RejectionNote,
		TotalDays:     l.TotalDays,
		LeaveType:     leaveType,
		CreatedAt:     l.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}

func (l *LeaveRequest) IsPending() bool {
	return l.Status == constant.PENDING
}
