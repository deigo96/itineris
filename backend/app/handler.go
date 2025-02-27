package app

import (
	"github.com/deigo96/itineris/app/config"
	"github.com/deigo96/itineris/app/internal/handler"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewHandler(config *config.Config, db *gorm.DB, router *gin.RouterGroup) {

	handler.NewEmployeeHandler(config, db, router)
	handler.NewAuthHandler(config, db, router)
}

func HandlePageNotFound(r *gin.Engine) {
	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(404, gin.H{
			"code":    4040,
			"message": "page not found",
		})
	})
}

func HandleNoMethod(r *gin.Engine) {
	r.NoMethod(func(ctx *gin.Context) {
		ctx.JSON(404, gin.H{
			"code":    4040,
			"message": "method not allowed",
		})
	})
}
