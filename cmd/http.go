package cmd

import (
	"ewallet-notification/helpers"
	"ewallet-notification/internal/api"
	"ewallet-notification/internal/services"
	"log"

	"github.com/gin-gonic/gin"
)

func ServeHTTP() {
	healthcheckSvc := &services.Healthcheck{}
	healthcheckAPI := &api.Healthcheck{
		HealthcheckServices: healthcheckSvc,
	}

	r := gin.Default()

	r.GET("/health", healthcheckAPI.HealthcheckHandlerHttp)

	err := r.Run(":" + helpers.GetEnv("PORT", "8080"))
	if err != nil {
		log.Fatal(err)
	}
}
