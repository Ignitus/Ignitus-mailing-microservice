package main

import (
	"os"

	"github.com/Marvin9/ignitus-mailing-microservice/api"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func usePort() string {
	port := os.Getenv("PORT")
	if port != "" {
		return ":" + port
	}
	return ":8000"
}

func main() {
	router := gin.Default()

	router.POST("/mail/confirmation", api.ConfirmationMailAPI)

	router.Run(usePort())
}
