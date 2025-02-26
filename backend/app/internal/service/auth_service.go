package service

import (
	"context"
	"errors"

	"github.com/deigo96/itineris/app/config"
	"github.com/deigo96/itineris/app/internal/entity"
	customError "github.com/deigo96/itineris/app/internal/error"
	"github.com/deigo96/itineris/app/internal/model"
	"github.com/deigo96/itineris/app/internal/repository"
	"github.com/deigo96/itineris/app/internal/util"
	"gorm.io/gorm"
)

type AuthService interface {
	Login(context.Context, *model.LoginRequest) (*model.LoginResponse, error)
}

type authService struct {
	db                 *gorm.DB
	config             *config.Config
	jwtService         JWTService
	employeeRepository repository.EmployeeRepository
}

func NewAuthService(db *gorm.DB, config *config.Config) AuthService {
	return &authService{
		db:                 db,
		config:             config,
		employeeRepository: repository.NewEmployeeRepository(),
		jwtService:         NewJWTService(config),
	}
}

func (s *authService) Login(c context.Context, req *model.LoginRequest) (*model.LoginResponse, error) {
	employee, err := s.employeeRepository.GetEmployeeByNip(c, s.db, req.NIP)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, customError.ErrIncorrectNipOrPassword
		}

		return nil, err
	}

	if !util.ValidatePassword(req.Password, employee.Password) {
		return nil, customError.ErrIncorrectNipOrPassword
	}

	token, err := s.generateToken(employee)
	if err != nil {
		return nil, err
	}

	return &model.LoginResponse{
		Token:     token,
		TokenType: "Bearer",
	}, nil
}

func (s *authService) generateToken(employee *entity.Employee) (string, error) {
	customClaims := model.CustomClaims{
		ID:   employee.ID,
		Nip:  employee.Nip,
		Role: employee.RoleId.String(),
	}

	token, err := s.jwtService.GenerateToken(customClaims)
	if err != nil {
		return "", err
	}

	return token, nil
}
