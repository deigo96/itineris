package service

import (
	"context"
	"errors"
	"strings"

	"github.com/deigo96/itineris/app/config"
	"github.com/deigo96/itineris/app/internal/entity"
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
	GetLeaveType(*gin.Context) ([]model.LeaveTypeResponse, error)
	CreateEmployee(c *gin.Context, req *model.CreateEmployeeRequest) (*model.EmployeeResponse, error)
}

type employeeService struct {
	EmployeeRepository     repository.EmployeeRepository
	db                     *gorm.DB
	config                 *config.Config
	repository             repository.Repository
	leaveRequestRepository repository.LeaveRequestRepository
}

func NewEmployeService(db *gorm.DB, config *config.Config) EmployeeService {
	return &employeeService{
		EmployeeRepository:     repository.NewEmployeeRepository(),
		db:                     db,
		config:                 config,
		repository:             repository.NewRepository(),
		leaveRequestRepository: repository.NewLeaveRequestRepository(),
	}
}

func (s *employeeService) CreateEmployee(c *gin.Context, req *model.CreateEmployeeRequest) (*model.EmployeeResponse, error) {
	user := util.GetContext(c)

	if err := s.validateCreateEmployee(user, req); err != nil {
		return nil, err
	}
	password, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	req.Password = password

	employeeEntity := &entity.Employee{
		CreatedBy: user.Nip,
		UpdatedBy: user.Nip,
	}

	employeeEntity.ToEntity(req)

	employee, err := s.EmployeeRepository.CreateEmployee(c, s.db, employeeEntity)
	if err != nil {
		return nil, err
	}

	return employee.ToModel(0, nil), nil
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
		totalPending := s.leaveRequestRepository.CountPendingRequest(c, s.db, user.ID, true)

		employeeResponse := user.ToModel(int(totalPending), nil)
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

	leaveRequest, err := s.leaveRequestRepository.GetLeaveRequests(c, s.db, ctx.IsAdmin(), employee.ID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	totalPending := s.leaveRequestRepository.CountPendingRequest(c, s.db, employee.ID, ctx.IsAdmin())

	return employee.ToModel(int(totalPending), leaveRequest), nil
}

func (s *employeeService) GetLeaveType(c *gin.Context) ([]model.LeaveTypeResponse, error) {
	ctx := util.GetContext(c)
	employee, err := s.EmployeeRepository.GetEmployeeByID(c, s.db, ctx.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, customError.ErrNotFound
		}
		return nil, err
	}

	res, err := s.repository.GetLeaveType(c, s.db, employee.IsPns)
	if err != nil {
		return nil, err
	}

	leaveTypeResponses := make([]model.LeaveTypeResponse, 0)
	for _, leaveType := range res {
		leaveTypeResponse := leaveType.ToModel()
		leaveTypeResponses = append(leaveTypeResponses, *leaveTypeResponse)
	}

	return leaveTypeResponses, nil
}

func (s *employeeService) validateCreateEmployee(c util.Context, req *model.CreateEmployeeRequest) error {
	if !c.IsAdmin() {
		return customError.ErrUnauthorized
	}

	employee, err := s.EmployeeRepository.GetEmployeeByNip(context.Background(), s.db, req.NIP)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if employee != nil {
		return customError.ErrNipAlreadyUsed
	}

	_, err = s.repository.GetRole(context.Background(), s.db, strings.ToUpper(req.Role))
	if err != nil {
		return customError.ErrInvalidRole
	}

	return nil
}
