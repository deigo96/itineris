package handler

import (
	"github.com/deigo96/itineris/app/config"
	"github.com/deigo96/itineris/app/internal/controller"
	"github.com/deigo96/itineris/app/internal/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewAuthHandler(config *config.Config, db *gorm.DB, router *gin.RouterGroup) {
	authService := service.NewAuthService(db, config)

	authController := controller.NewAuthController(authService)
	auth := router.Group("/auth")

	auth.POST("/login", authController.Login)
}
