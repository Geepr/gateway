package main

import (
	"github.com/Geepr/gateway/controllers"
	"github.com/Geepr/gateway/services"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true, DisableColors: true})
	log.SetOutput(os.Stdout)
	//todo: configurable
	log.SetLevel(log.InfoLevel)
}

func main() {
	router := setupEngine()

	//todo: configurable address
	if err := router.Run("localhost:5510"); err != nil {
		log.Panicf("Server failed while listening: %s", err.Error())
	}
}

func setupEngine() *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(services.GetGinLogger())
	//todo: trusted proxies

	//todo: base path config
	basePath := ""
	controllers.SetupGameRoutes(router, basePath)

	return router
}
