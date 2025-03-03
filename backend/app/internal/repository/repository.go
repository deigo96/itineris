package repository

import (
	"context"

	"github.com/deigo96/itineris/app/internal/entity"
	"gorm.io/gorm"
)

type Repository interface {
	GetLeaveType(c context.Context, db *gorm.DB, isPns bool) ([]entity.LeaveType, error)
	GetLeaveTypeByID(c context.Context, db *gorm.DB, id int) (*entity.LeaveType, error)
	GetRole(c context.Context, db *gorm.DB, role string) (*entity.Role, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetLeaveType(c context.Context, db *gorm.DB, isPns bool) ([]entity.LeaveType, error) {
	var leaveType []entity.LeaveType

	query := db.Order("id ASC")
	if isPns {
		query.Where("is_pns = ?", true)
	} else {
		query.Where("is_pppk = ?", true)
	}

	if err := query.Find(&leaveType).Error; err != nil {
		return nil, err
	}

	return leaveType, nil
}

func (r *repository) GetLeaveTypeByID(c context.Context, db *gorm.DB, id int) (*entity.LeaveType, error) {
	var res entity.LeaveType
	if err := db.Where("id = ?", id).First(&res).Error; err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *repository) GetRole(c context.Context, db *gorm.DB, role string) (*entity.Role, error) {
	var res entity.Role
	if err := db.Where("role_name = ?", role).First(&res).Error; err != nil {
		return nil, err
	}

	return &res, nil
}
