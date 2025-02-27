package controller

import (
	"fmt"

	customError "github.com/deigo96/itineris/app/internal/error"
	"github.com/deigo96/itineris/app/internal/model"
	"github.com/deigo96/itineris/app/internal/service"
	"github.com/deigo96/itineris/app/internal/util"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type LeaveRequestController struct {
	LeaveRequestService service.LeaveRequestService
}

func NewLeaveRequestController(leaveRequestService service.LeaveRequestService) LeaveRequestController {
	return LeaveRequestController{LeaveRequestService: leaveRequestService}
}

func (c *LeaveRequestController) LeaveRequest(ctx *gin.Context) {
	var req model.LeaveRequestRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("error: ", err)
		customError.ErrorResponse(err, ctx)
		return
	}

	validate = validator.New()
	if err := validate.Struct(req); err != nil {
		customError.ErrorResponse(err, ctx)
		return
	}

	err := c.LeaveRequestService.LeaveRequest(ctx, &req)
	if err != nil {
		customError.ErrorResponse(err, ctx)
		return
	}

	ctx.JSON(200, util.SuccessResponse(nil))

}
