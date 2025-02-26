package app

import (
	"github.com/deigo96/itineris/app/config"
	"github.com/deigo96/itineris/app/internal/handler"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewHandler(config *config.Config, db *gorm.DB, router *gin.RouterGroup) {

	handler.NewUserHandler(config, db, router)
}
