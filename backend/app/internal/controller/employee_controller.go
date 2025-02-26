package controller

import (
	customError "github.com/deigo96/itineris/app/internal/error"
	"github.com/deigo96/itineris/app/internal/util"

	"github.com/deigo96/itineris/app/internal/service"
	"github.com/gin-gonic/gin"
)

type EmployeeController struct {
	EmployeeService service.EmployeeService
}

func NewEmployeeController(userService service.EmployeeService) EmployeeController {
	return EmployeeController{EmployeeService: userService}
}

func (c *EmployeeController) Get(ctx *gin.Context) {
	employee, err := c.EmployeeService.GetEmployee(ctx)
	if err != nil {
		customError.ErrorResponse(err, ctx)
		return
	}
	ctx.JSON(200, util.SuccessResponse(employee))
}
