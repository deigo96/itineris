package handler

import (
	"github.com/deigo96/itineris/app/config"
	"github.com/deigo96/itineris/app/internal/controller"
	"github.com/deigo96/itineris/app/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewEmployeeHandler(config *config.Config, db *gorm.DB, router *gin.RouterGroup) {
	employeeService := service.NewEmployeService(db, config)

	employeeController := controller.NewEmployeeController(employeeService)
	userRoute := router.Group("/users")

	userRoute.GET("", employeeController.Get)

}
