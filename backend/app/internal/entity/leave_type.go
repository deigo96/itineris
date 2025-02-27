package entity

import "github.com/deigo96/itineris/app/internal/model"

type LeaveType struct {
	ID       int `gorm:"primaryKey"`
	TypeName string
	IsPns    bool
	IsPppk   bool
}

func (l *LeaveType) ToModel() *model.LeaveTypeResponse {
	return &model.LeaveTypeResponse{
		ID: l.ID, TypeName: l.TypeName,
	}
}
