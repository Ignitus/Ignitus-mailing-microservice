package main

import (
	"os"

	"github.com/Ignitus/ignitus-mailing-microservice/api"
	"github.com/Ignitus/ignitus-mailing-microservice/api/middleware"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

var acceptedHost = os.Getenv("ACCEPTED_HOST")

func usePort() string {
	port := os.Getenv("PORT")
	if port != "" {
		return ":" + port
	}
	return ":8000"
}

func main() {
	router := gin.Default()

	router.Use(middleware.APIAccessTo(acceptedHost))

	router.POST("/mail/confirmation", api.ConfirmationMailAPI)

	router.Run(usePort())
}
