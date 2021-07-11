package router

import (
	"github.com/gin-gonic/gin"
	"log"
)

func Init() {
	router := gin.Default()
	router.Group("/")

	err := router.Run(":80801")
	if err != nil {
		log.Fatalf("failed to start the server: %v", err)
	}
}