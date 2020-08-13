package main

import (
	"github.com/Marvin9/ignitus-mailing-microservice/api"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	router := gin.Default()

	router.POST("/mail/confirmation", api.ConfirmationMailAPI)

	router.Run(":8000")
}
