package controllers

import (
	"fmt"
	"github.com/Geepr/gateway/clients/game_client"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func GetPlatforms(c *gin.Context) {
	var query struct {
		PageIndex int    `form:"page" binding:"required"`
		PageSize  int    `form:"pageSize" binding:"required"`
		Name      string `form:"name"`
	}
	if err := c.BindQuery(&query); err != nil {
		log.Infof("Failed to bind platform query: %s", err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	platforms, _, err := game_client.GetPlatforms(query.PageIndex, query.PageSize, query.Name)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, platforms)
}

func SetupPlatformRoutes(engine *gin.Engine, basePath string) {
	baseUrl := fmt.Sprintf("%s/api/v0/platforms", basePath)

	engine.GET(baseUrl, GetPlatforms)
}
