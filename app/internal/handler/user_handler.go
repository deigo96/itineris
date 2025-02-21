package handler

import (
	"github.com/deigo96/bpkp/app/config"
	"github.com/deigo96/bpkp/app/internal/controller"
	"github.com/deigo96/bpkp/app/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewUserHandler(config *config.Config, db *gorm.DB, router *gin.RouterGroup) {
	userService := service.NewUserService(db, config)

	userController := controller.NewUserController(userService)
	userRoute := router.Group("/users")

	userRoute.GET("", userController.Get)

}
