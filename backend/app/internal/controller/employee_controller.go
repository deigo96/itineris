package controller

import (
	"net/http"

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
	employees, err := c.EmployeeService.GetEmployees(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}
	ctx.JSON(200, employees)
}
