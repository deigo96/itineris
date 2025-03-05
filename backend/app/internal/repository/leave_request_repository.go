package repository

import (
	"context"

	constant "github.com/deigo96/itineris/app/internal/const"
	"github.com/deigo96/itineris/app/internal/entity"
	"gorm.io/gorm"
)

type LeaveRequestRepository interface {
	StoreLeaveRequest(context.Context, *gorm.DB, *entity.LeaveRequest) (*entity.LeaveRequest, error)
	UpdateLeaveRequest(context.Context, *gorm.DB, int, *entity.UpdateLeaveRequest) error
	GetLeaveRequestByID(context.Context, *gorm.DB, int, int, bool) (*entity.LeaveRequest, error)
	GetLeaveRequests(context.Context, *gorm.DB, bool, int) ([]*entity.LeaveRequest, error)
	GetRequestByStatus(context.Context, *gorm.DB, int, bool, constant.Status) ([]*entity.LeaveRequest, error)
	CountPendingRequest(
		c context.Context, db *gorm.DB, employeeID int, isAdmin bool) int64
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

func (r *leaveRequestRepository) GetRequestByStatus(
	c context.Context, db *gorm.DB, employeeID int, isAdmin bool, status constant.Status) ([]*entity.LeaveRequest, error) {
	var leaveRequests []*entity.LeaveRequest

	if !isAdmin {
		if err := db.Where("employee_id = ? AND status = ?", employeeID, status).Order("id DESC").Find(&leaveRequests).Error; err != nil {
			return nil, err
		}
		return leaveRequests, nil
	}

	if err := db.Where("status = ?", status).Order("id DESC").Find(&leaveRequests).Error; err != nil {
		return nil, err
	}

	return leaveRequests, nil
}

func (r *leaveRequestRepository) CountPendingRequest(
	c context.Context, db *gorm.DB, employeeID int, isAdmin bool) int64 {

	var total int64
	if !isAdmin {
		db.Model(&entity.LeaveRequest{}).Where("employee_id = ? AND status = ?", employeeID, constant.PENDING).Count(&total)
	} else {
		db.Model(&entity.LeaveRequest{}).Where("status = ?", constant.PENDING).Count(&total)
	}

	return total
}
