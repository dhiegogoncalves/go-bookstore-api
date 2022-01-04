package app

import (
	"bookstore-users-api/logger"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	mapUrls()
	logger.Info("Starting application")
	router.Run(":8080")
}
