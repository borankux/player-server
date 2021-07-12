package router

import (
	"fmt"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/player-server/conf"
	"log"
)

func Init() {
	port := conf.GetConf().GetString("app.port")
	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("player/spa/dist", false)))
	err := router.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to start the server: %v", err)
	}
}