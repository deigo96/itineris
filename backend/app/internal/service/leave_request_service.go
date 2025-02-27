package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/deigo96/itineris/app/config"
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
}

type leaveRequestService struct {
	db                     *gorm.DB
	config                 *config.Config
	leaveRequestRepository repository.LeaveRequestRepository
	employeeRepository     repository.EmployeeRepository
}

func NewLeaveRequestService(db *gorm.DB, config *config.Config) LeaveRequestService {
	return &leaveRequestService{
		db:                     db,
		config:                 config,
		leaveRequestRepository: repository.NewLeaveRequestRepository(),
		employeeRepository:     repository.NewEmployeeRepository(),
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
