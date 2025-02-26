package service

import (
	"context"

	"github.com/deigo96/itineris/app/config"
	"github.com/deigo96/itineris/app/internal/model"
	"github.com/deigo96/itineris/app/internal/repository"
	"gorm.io/gorm"
)

type EmployeeService interface {
	GetEmployees(context.Context) ([]model.EmployeeResponse, error)
}

type employeeService struct {
	EmployeeRepository repository.EmployeeRepository
	db                 *gorm.DB
	config             *config.Config
}

func NewEmployeService(db *gorm.DB, config *config.Config) EmployeeService {
	return &employeeService{
		EmployeeRepository: repository.NewEmployeeRepository(),
		db:                 db,
		config:             config,
	}
}

func (s *employeeService) GetEmployees(c context.Context) ([]model.EmployeeResponse, error) {
	users, err := s.EmployeeRepository.GetEmployees(c, s.db)
	if err != nil {
		return nil, err
	}

	employeeResponses := make([]model.EmployeeResponse, 0)
	if users == nil {
		return employeeResponses, nil
	}

	for _, user := range users {
		employeeResponse := user.ToModel()
		employeeResponses = append(employeeResponses, *employeeResponse)
	}

	return employeeResponses, nil
}

// func (s *employeeService) CreateUser(c context.Context, user *model.CreateUserRequest) (*model.EmployeeResponse, error) {

// }

// func (s *employeeService) constructCreateUser(c context.Context, user *model.CreateUserRequest) (*model.Users, error) {

// }
