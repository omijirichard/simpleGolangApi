package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(currentContext *gin.Context) {
		currentContext.JSON(http.StatusOK, gin.H{"message": "API is running"})
	})

	router.GET("/health", func(currentContext *gin.Context) {
		currentContext.JSON(http.StatusOK, gin.H{"message": "healthy"})
	})

	router.GET("/me", func(currentContext *gin.Context) {
		currentContext.JSON(http.StatusOK, gin.H{"name": "Omiji Richard", "email": "omijirichard@gmail.com", "github": "https://github.com/omijirichard"})
	})
	router.Run(":2222")
}
