package repository

import (
	"context"

	"github.com/deigo96/itineris/app/internal/entity"
	"gorm.io/gorm"
)

type EmployeeRepository interface {
	GetEmployees(context.Context, *gorm.DB) ([]entity.Employee, error)
	GetEmployeeByNip(context.Context, *gorm.DB, string) (*entity.Employee, error)
	GetEmployeeByID(context.Context, *gorm.DB, int) (*entity.Employee, error)
	CreateEmployee(context.Context, *gorm.DB, *entity.Employee) (*entity.Employee, error)
	UpdateBalance(context.Context, *gorm.DB, int, int) error
	RestoreBalance(context.Context, *gorm.DB, int, int) error
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

func (r *employeeRepository) GetEmployeeByNip(c context.Context, db *gorm.DB, nip string) (*entity.Employee, error) {
	employee := &entity.Employee{}
	if err := db.Where("nip = ?", nip).First(&employee).Error; err != nil {
		return nil, err
	}

	return employee, nil
}

func (r *employeeRepository) GetEmployeeByID(c context.Context, db *gorm.DB, id int) (*entity.Employee, error) {
	employee := &entity.Employee{}
	if err := db.Where("id = ?", id).First(&employee).Error; err != nil {
		return nil, err
	}

	return employee, nil
}

func (r *employeeRepository) CreateEmployee(c context.Context, db *gorm.DB, employee *entity.Employee) (*entity.Employee, error) {
	if err := db.Create(&employee).Error; err != nil {
		return nil, err
	}

	return employee, nil
}

func (r *employeeRepository) UpdateBalance(c context.Context, db *gorm.DB, totalRequest, ID int) error {
	query := `
	UPDATE employees
	SET leave_balance = leave_balance - ?
	WHERE id = ?
	`
	if err := db.Exec(query, totalRequest, ID).Error; err != nil {
		return err
	}

	return nil
}

func (r *employeeRepository) RestoreBalance(c context.Context, db *gorm.DB, totalRequest, ID int) error {
	query := `
	UPDATE employees
	SET leave_balance = leave_balance + ?
	WHERE id = ?
	`
	if err := db.Exec(query, totalRequest, ID).Error; err != nil {
		return err
	}

	return nil
}
