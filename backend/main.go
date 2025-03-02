package main

import (
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

	app.HandlePageNotFound(r)
	app.HandleNoMethod(r)
	v1 := r.Group("/api/v1")

	app.NewHandler(configuration, db, v1)

	log.Println("Service " + configuration.ServiceName + " running on port " + configuration.ServicePort)
	r.Run(":" + configuration.ServicePort)

	config.DBCloseConnection(db)
}
