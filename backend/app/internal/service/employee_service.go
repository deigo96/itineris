package service

import (
	"context"
	"errors"

	"github.com/deigo96/itineris/app/config"
	customError "github.com/deigo96/itineris/app/internal/error"
	"github.com/deigo96/itineris/app/internal/model"
	"github.com/deigo96/itineris/app/internal/repository"
	"github.com/deigo96/itineris/app/internal/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type EmployeeService interface {
	GetEmployees(context.Context) ([]model.EmployeeResponse, error)
	GetEmployee(*gin.Context) (*model.EmployeeResponse, error)
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

func (s *employeeService) GetEmployee(c *gin.Context) (*model.EmployeeResponse, error) {
	ctx := util.GetContext(c)

	employee, err := s.EmployeeRepository.GetEmployeeByID(c, s.db, ctx.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, customError.ErrNotFound
		}
		return nil, err
	}

	return employee.ToModel(), nil
}
