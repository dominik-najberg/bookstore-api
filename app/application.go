package app

import (
	"github.com/gin-gonic/gin"
	"log"
)

var (
	router = gin.Default()
)

func StartApplication()  {
	mapUrls()
	if err := router.Run("localhost:8080"); err != nil {
		log.Fatalf("error while runing router: %v", err)
	}
}