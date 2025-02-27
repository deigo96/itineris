package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/deigo96/itineris/app/config"
	constant "github.com/deigo96/itineris/app/internal/const"
	"github.com/deigo96/itineris/app/internal/entity"
	customError "github.com/deigo96/itineris/app/internal/error"
	"github.com/deigo96/itineris/app/internal/model"
	"github.com/deigo96/itineris/app/internal/repository"
	"github.com/deigo96/itineris/app/internal/util"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LeaveRequestService interface {
	LeaveRequest(c *gin.Context, req *model.LeaveRequestRequest) error
	Approval(c *gin.Context, req *model.ApprovalRequest) error
	GetLeaveRequests(c *gin.Context) ([]model.LeaveRequestResponse, error)
}

type leaveRequestService struct {
	db                     *gorm.DB
	config                 *config.Config
	leaveRequestRepository repository.LeaveRequestRepository
	employeeRepository     repository.EmployeeRepository
	repository             repository.Repository
}

func NewLeaveRequestService(db *gorm.DB, config *config.Config) LeaveRequestService {
	return &leaveRequestService{
		db:                     db,
		config:                 config,
		leaveRequestRepository: repository.NewLeaveRequestRepository(),
		employeeRepository:     repository.NewEmployeeRepository(),
		repository:             repository.NewRepository(),
	}
}

func (s *leaveRequestService) LeaveRequest(c *gin.Context, req *model.LeaveRequestRequest) error {
	ctx := util.GetContext(c)

	employee, err := s.employeeRepository.GetEmployeeByID(c, s.db, ctx.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return customError.ErrNotFound
		}

		return err
	}

	totalRequest, err := s.validateLeaveRequest(employee, req)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	reqEntity := &entity.LeaveRequest{}
	reqEntity.ToEntity(req)
	reqEntity.EmployeeID = ctx.ID

	return s.process(c.Request.Context(), totalRequest, reqEntity)
}

func (s *leaveRequestService) process(c context.Context, totalRequest int, req *entity.LeaveRequest) (err error) {
	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
			err = r.(error)
			tx.Rollback()
		}
	}()

	if err := s.employeeRepository.UpdateBalance(c, tx, totalRequest, req.EmployeeID); err != nil {
		log.Println("error update balance: ", err)
		tx.Rollback()
		return err
	}

	req.TotalDays = totalRequest
	_, err = s.leaveRequestRepository.StoreLeaveRequest(c, tx, req)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err.Error
	}
	return nil
}

func (s *leaveRequestService) validateLeaveRequest(
	employee *entity.Employee, req *model.LeaveRequestRequest) (int, error) {
	if employee.LeaveBalance <= 0 {
		return 0, customError.ErrLeaveBalance
	}

	startDate, err := util.ParseStringToTime(req.StartDate)
	if err != nil {
		return 0, err
	}
	endDate, err := util.ParseStringToTime(req.EndDate)
	if err != nil {
		return 0, err
	}

	if startDate.Before(time.Now()) || endDate.Before(time.Now()) || startDate.After(endDate) {
		return 0, customError.ErrTimeLeaveRequest
	}

	totalRequest := endDate.Sub(startDate).Hours() / 24
	if employee.LeaveBalance < int32(totalRequest) {
		return 0, customError.ErrLeaveBalance
	}

	return int(totalRequest + 1), nil
}

func (s *leaveRequestService) Approval(c *gin.Context, req *model.ApprovalRequest) error {
	ctx := util.GetContext(c)

	employee, err := s.employeeRepository.GetEmployeeByID(c, s.db, ctx.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return customError.ErrNotFound
		}

		return err
	}

	var isAdmin bool = true
	if employee.RoleId != constant.PPK {
		return customError.ErrUnauthorized
	}

	leaveRequest, err := s.leaveRequestRepository.GetLeaveRequestByID(c, s.db, employee.ID, req.ID, isAdmin)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return customError.ErrNotFound
		}

		return err
	}

	if !leaveRequest.IsPending() {
		return customError.ErrLeaveRequestHasBeenProcessed
	}

	switch req.ApprovalStatus {
	case constant.APPROVE:
		return s.approve(c, req)
	case constant.REJECT:
		return s.reject(c, leaveRequest, req)
	default:
		return customError.ErrInvalidApprovalStatus
	}

}

func (s *leaveRequestService) approve(c *gin.Context, req *model.ApprovalRequest) error {
	user := util.GetContext(c)
	approvalRequest := &entity.UpdateLeaveRequest{
		Status:    constant.Status(req.ApprovalStatus.String()),
		UpdatedBy: user.Nip,
	}

	return s.leaveRequestRepository.UpdateLeaveRequest(c, s.db, req.ID, approvalRequest)
}

func (s *leaveRequestService) reject(c *gin.Context, leaveRequest *entity.LeaveRequest, req *model.ApprovalRequest) (err error) {
	user := util.GetContext(c)

	approvalRequest := &entity.UpdateLeaveRequest{
		Status:        constant.Status(req.ApprovalStatus.String()),
		UpdatedBy:     user.Nip,
		RejectionNote: req.RejectionNote,
	}

	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
			err = r.(error)
			tx.Rollback()
		}
	}()

	if err := s.leaveRequestRepository.UpdateLeaveRequest(c, tx, req.ID, approvalRequest); err != nil {
		tx.Rollback()
		return err
	}

	if err := s.employeeRepository.RestoreBalance(c, tx, int(leaveRequest.TotalDays), leaveRequest.EmployeeID); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (s *leaveRequestService) GetLeaveRequests(c *gin.Context) ([]model.LeaveRequestResponse, error) {
	user := util.GetContext(c)

	responses, err := s.leaveRequestRepository.GetLeaveRequests(c, s.db, user.IsAdmin(), user.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, customError.ErrNotFound
		}
		return nil, err
	}

	leaveRequestResponse := make([]model.LeaveRequestResponse, 0)

	for _, response := range responses {
		leaveType, err := s.repository.GetLeaveTypeByID(c, s.db, response.LeaveType)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, customError.ErrNotFound
			}
			return nil, err
		}
		res := response.ToModel(leaveType.TypeName)

		leaveRequestResponse = append(leaveRequestResponse, *res)
	}

	return leaveRequestResponse, nil
}
