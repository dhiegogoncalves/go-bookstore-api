package app

import (
	"bookstore-oauth-api/src/domain/access_token"
	"bookstore-oauth-api/src/http"
	"bookstore-oauth-api/src/repository/db"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	atHandler := http.NewHandler(access_token.NewService(db.NewRepository()))

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)

	router.Run(":8080")
}
