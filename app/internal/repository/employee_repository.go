package repository

import (
	"context"

	"github.com/deigo96/itineris/app/internal/entity"
	"gorm.io/gorm"
)

type EmployeeRepository interface {
	GetEmployees(context.Context, *gorm.DB) ([]entity.Employee, error)
	CreateUser(context.Context, *gorm.DB, *entity.Employee) (*entity.Employee, error)
}

type employeeRepository struct{}

func NewEmployeeRepository() EmployeeRepository {
	return &employeeRepository{}
}

func (r *employeeRepository) GetEmployees(c context.Context, db *gorm.DB) ([]entity.Employee, error) {
	employees := []entity.Employee{}
	if err := db.Find(&employees).Error; err != nil {
		return nil, err
	}

	return employees, nil
}

func (r *employeeRepository) CreateUser(c context.Context, db *gorm.DB, employee *entity.Employee) (*entity.Employee, error) {
	if err := db.Create(&employee).Error; err != nil {
		return nil, err
	}

	return employee, nil
}
