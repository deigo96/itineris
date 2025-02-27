package model

type LeaveRequestRequest struct {
	StartDate string `json:"start_date" time_format:"2006-01-02" validate:"required"`
	EndDate   string `json:"end_date" time_format:"2006-01-02" validate:"required"`
	Reason    string `json:"reason" validate:"required"`
	LeaveType int    `json:"leave_type" validate:"required"`
}
