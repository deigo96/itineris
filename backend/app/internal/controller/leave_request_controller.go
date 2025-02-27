package controller

import (
	"fmt"
	"strconv"

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

func (c *LeaveRequestController) Action(ctx *gin.Context) {
	var req model.ApprovalRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		customError.ErrorResponse(err, ctx)
		return
	}

	validate = validator.New()
	if err := validate.Struct(req); err != nil {
		customError.ErrorResponse(err, ctx)
		return
	}

	err := c.LeaveRequestService.Approval(ctx, &req)
	if err != nil {
		customError.ErrorResponse(err, ctx)
		return
	}

	ctx.JSON(200, util.SuccessResponse(nil))
}

func (c *LeaveRequestController) GetLeaveRequests(ctx *gin.Context) {
	responses, err := c.LeaveRequestService.GetLeaveRequests(ctx)
	if err != nil {
		customError.ErrorResponse(err, ctx)
		return
	}

	ctx.JSON(200, util.SuccessResponse(responses))
}

func (c *LeaveRequestController) GetLeaveRequest(ctx *gin.Context) {
	paramID := ctx.Param("id")
	id, err := strconv.Atoi(paramID)
	if err != nil || id == 0 {
		customError.ErrorResponse(customError.ErrBadRequest, ctx)
		return
	}

	response, err := c.LeaveRequestService.GetLeaveRequest(ctx, id)
	if err != nil {
		customError.ErrorResponse(err, ctx)
		return
	}

	ctx.JSON(200, util.SuccessResponse(response))
}
