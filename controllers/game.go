package controllers

import (
	"fmt"
	"github.com/Geepr/gateway/clients/game_client"
	"github.com/Geepr/gateway/clients/game_client/dto"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func GetGames(c *gin.Context) {
	var query struct {
		PageIndex int    `form:"page" binding:"required"`
		PageSize  int    `form:"pageSize" binding:"required"`
		Title     string `form:"title"`
	}
	if err := c.BindQuery(&query); err != nil {
		log.Infof("Failed to bind game query: %s", err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	games, _, err := game_client.GetGames(query.PageIndex, query.PageSize, query.Title)
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

func UpdateGame(c *gin.Context) {
	var updateDto dto.GameUpdateDto
	if err := c.BindJSON(&updateDto); err != nil {
		log.Infof("Failed to parse game update model: %s", err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	id, err := parseUuidFromParam(c)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	responseCode, err := game_client.UpdateGame(id, &updateDto)
	if err != nil {
		if responseCode == http.StatusNotFound {
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			c.AbortWithStatus(http.StatusInternalServerError)
		}
		return
	}
	c.Status(http.StatusOK)
}

func DeleteGame(c *gin.Context) {
	lookupUuid, err := parseUuidFromParam(c)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	responseCode, err := game_client.DeleteGame(lookupUuid)
	if err != nil {
		if responseCode == http.StatusNotFound {
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			c.AbortWithStatus(http.StatusInternalServerError)
		}
		return
	}

	c.Status(http.StatusOK)
}

func CreateGame(c *gin.Context) {
	var createDto dto.GameCreateDto
	if err := c.BindJSON(&createDto); err != nil {
		log.Infof("Failed to parse game create model: %s", err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	responseBody, responseCode, err := game_client.CreateGame(&createDto)
	if err != nil {
		if responseCode == http.StatusNotFound {
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			c.AbortWithStatus(http.StatusInternalServerError)
		}
		return
	}

	c.JSON(http.StatusCreated, responseBody)
}

func SetupGameRoutes(engine *gin.Engine, basePath string) {
	baseUrl := fmt.Sprintf("%s/api/v0/games", basePath)

	engine.GET(baseUrl+"/:id", GetGame)
	engine.GET(baseUrl, GetGames)
	engine.POST(baseUrl+"/:id", UpdateGame)
	engine.POST(baseUrl, CreateGame)
	engine.DELETE(baseUrl+"/:id", DeleteGame)
}
