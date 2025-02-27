package model

import (
	constant "github.com/deigo96/itineris/app/internal/const"
)

type LeaveRequestRequest struct {
	StartDate string `json:"start_date" time_format:"2006-01-02" validate:"required"`
	EndDate   string `json:"end_date" time_format:"2006-01-02" validate:"required"`
	Reason    string `json:"reason" validate:"required"`
	LeaveType int    `json:"leave_type" validate:"required"`
}

type LeaveRequestResponse struct {
	ID            int    `json:"id"`
	EmployeeID    int    `json:"employee_id"`
	Status        string `json:"status"`
	StartDate     string `json:"start_date"`
	EndDate       string `json:"end_date"`
	Reason        string `json:"reason"`
	RejectionNote string `json:"rejection_note"`
	TotalDays     int    `json:"total_days"`
	LeaveType     string `json:"leave_type"`
	CreatedAt     string `json:"created_at"`
}

type ApprovalRequest struct {
	ID             int               `json:"id" validate:"required"`
	ApprovalStatus constant.Approval `json:"status" validate:"required"`
	RejectionNote  string            `json:"rejection_note"`
}
