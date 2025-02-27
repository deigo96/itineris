package repository

import (
	"context"

	"github.com/deigo96/itineris/app/internal/entity"
	"gorm.io/gorm"
)

type LeaveRequestRepository interface {
	StoreLeaveRequest(context.Context, *gorm.DB, *entity.LeaveRequest) (*entity.LeaveRequest, error)
}

type leaveRequestRepository struct{}

func NewLeaveRequestRepository() LeaveRequestRepository {
	return &leaveRequestRepository{}
}

func (r *leaveRequestRepository) StoreLeaveRequest(c context.Context,
	db *gorm.DB, leaveRequest *entity.LeaveRequest) (
	*entity.LeaveRequest, error) {

	if err := db.Create(&leaveRequest).Error; err != nil {
		return nil, err
	}

	return leaveRequest, nil
}
