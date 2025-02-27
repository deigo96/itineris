package repository

import (
	"context"

	"github.com/deigo96/itineris/app/internal/entity"
	"gorm.io/gorm"
)

type Repository interface {
	GetLeaveType(c context.Context, db *gorm.DB, isPns bool) ([]entity.LeaveType, error)
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
