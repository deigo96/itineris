package main

import (
	"fmt"
	"log"

	"github.com/deigo96/itineris/app"
	"github.com/deigo96/itineris/app/config"
	"github.com/gin-gonic/gin"
)

func main() {

	configuration := config.GetConfig()
	db := config.DBConnection(configuration)

	if configuration.Environment == "PRODUCTION" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		fmt.Println("🚀 Received Request:", c.Request.Method, c.Request.URL.Path)
		c.Next()
	})

	app.HandlePageNotFound(r)
	app.HandleNoMethod(r)
	v1 := r.Group("/api/v1")

	app.NewHandler(configuration, db, v1)

	log.Println("Service " + configuration.ServiceName + " running on port " + configuration.ServicePort)
	r.Run(":" + configuration.ServicePort)

	config.DBCloseConnection(db)
}
