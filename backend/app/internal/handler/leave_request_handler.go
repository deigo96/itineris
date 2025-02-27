package handler

import (
	"github.com/deigo96/itineris/app/config"
	"github.com/deigo96/itineris/app/internal/controller"
	"github.com/deigo96/itineris/app/internal/middleware"
	"github.com/deigo96/itineris/app/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewLeaveRequestHandler(config *config.Config, db *gorm.DB, router *gin.RouterGroup) {
	leaveRequestService := service.NewLeaveRequestService(db, config)

	leaveRequestController := controller.NewLeaveRequestController(leaveRequestService)
	leaveRequest := router.Group("/leave-requests")
	leaveRequest.Use(middleware.Authorization(config))

	leaveRequest.POST("", leaveRequestController.LeaveRequest)
}
