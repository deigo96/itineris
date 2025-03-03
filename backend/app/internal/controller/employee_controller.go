package controller

import (
	customError "github.com/deigo96/itineris/app/internal/error"
	"github.com/deigo96/itineris/app/internal/model"
	"github.com/deigo96/itineris/app/internal/util"
	"github.com/go-playground/validator/v10"

	"github.com/deigo96/itineris/app/internal/service"
	"github.com/gin-gonic/gin"
)

type EmployeeController struct {
	EmployeeService service.EmployeeService
}

func NewEmployeeController(userService service.EmployeeService) EmployeeController {
	return EmployeeController{EmployeeService: userService}
}

func (c *EmployeeController) CreateEmployee(ctx *gin.Context) {
	req := &model.CreateEmployeeRequest{}
	if err := ctx.ShouldBindJSON(req); err != nil {
		customError.ErrorResponse(err, ctx)
		return
	}

	validate = validator.New()
	if err := validate.Struct(req); err != nil {
		customError.ErrorResponse(err, ctx)
		return
	}

	employee, err := c.EmployeeService.CreateEmployee(ctx, req)
	if err != nil {
		customError.ErrorResponse(err, ctx)
		return
	}

	ctx.JSON(200, util.SuccessResponse(employee))
}

func (c *EmployeeController) Get(ctx *gin.Context) {
	employee, err := c.EmployeeService.GetEmployee(ctx)
	if err != nil {
		customError.ErrorResponse(err, ctx)
		return
	}
	ctx.JSON(200, util.SuccessResponse(employee))
}

func (c *EmployeeController) GetLeaveType(ctx *gin.Context) {
	res, err := c.EmployeeService.GetLeaveType(ctx)
	if err != nil {
		customError.ErrorResponse(err, ctx)
		return
	}

	ctx.JSON(200, util.SuccessResponse(res))
}
