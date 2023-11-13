package controllers

import (
	"fmt"
	"github.com/Geepr/gateway/clients/game_client"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func GetGames(c *gin.Context) {
	var query struct {
		PageIndex int `form:"page" binding:"required"`
		PageSize  int `form:"pageSize" binding:"required"`
	}
	if err := c.BindQuery(&query); err != nil {
		log.Infof("Failed to bind game query: %s", err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	games, _, err := game_client.GetGames(query.PageIndex, query.PageSize)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, games)
}

func GetGame(c *gin.Context) {
	lookupUuid, err := parseUuidFromParam(c)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	game, responseCode, err := game_client.GetGame(lookupUuid)
	if err != nil {
		if responseCode == http.StatusNotFound {
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			c.AbortWithStatus(http.StatusInternalServerError)
		}
		return
	}

	c.JSON(http.StatusOK, game)
}

func SetupGameRoutes(engine *gin.Engine, basePath string) {
	baseUrl := fmt.Sprintf("%s/api/v0/games", basePath)

	engine.GET(baseUrl+"/:id", GetGame)
	engine.GET(baseUrl, GetGames)
}
