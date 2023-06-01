package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Working")
	})

	router.Run("localhost:8081")
}
