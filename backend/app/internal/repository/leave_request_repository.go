package repository

import (
	"context"

	"github.com/deigo96/itineris/app/internal/entity"
	"gorm.io/gorm"
)

type LeaveRequestRepository interface {
	StoreLeaveRequest(context.Context, *gorm.DB, *entity.LeaveRequest) (*entity.LeaveRequest, error)
	UpdateLeaveRequest(context.Context, *gorm.DB, int, *entity.UpdateLeaveRequest) error
	GetLeaveRequestByID(context.Context, *gorm.DB, int, int, bool) (*entity.LeaveRequest, error)
	GetLeaveRequests(context.Context, *gorm.DB, bool, int) ([]*entity.LeaveRequest, error)
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

func (r *leaveRequestRepository) UpdateLeaveRequest(c context.Context, db *gorm.DB, id int, req *entity.UpdateLeaveRequest) error {
	if err := db.Model(&entity.LeaveRequest{}).
		Where("id = ?", id).Updates(req).Error; err != nil {
		return err
	}
	return nil
}

func (r *leaveRequestRepository) GetLeaveRequestByID(
	c context.Context, db *gorm.DB, employeeID, id int, isAdmin bool) (*entity.LeaveRequest, error) {
	var leaveRequest entity.LeaveRequest

	if !isAdmin {
		if err := db.Where("employee_id = ?", employeeID).Where("id = ?", id).First(&leaveRequest).Error; err != nil {
			return nil, err
		}
		return &leaveRequest, nil
	}

	if err := db.Where("id = ?", id).First(&leaveRequest).Error; err != nil {
		return nil, err
	}
	return &leaveRequest, nil
}

func (r *leaveRequestRepository) GetLeaveRequests(
	c context.Context, db *gorm.DB, isAdmin bool, employeeID int) ([]*entity.LeaveRequest, error) {
	var leaveRequests []*entity.LeaveRequest

	if !isAdmin {
		if err := db.Where("employee_id = ?", employeeID).Order("id DESC").Find(&leaveRequests).Error; err != nil {
			return nil, err
		}
		return leaveRequests, nil
	}

	if err := db.Order("id DESC").Find(&leaveRequests).Error; err != nil {
		return nil, err
	}
	return leaveRequests, nil

}
