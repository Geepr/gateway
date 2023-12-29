package controllers

import (
	"fmt"
	"github.com/Geepr/gateway/clients/game_client"
	"github.com/Geepr/gateway/clients/game_client/dto"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func GetPlatform(c *gin.Context) {
	lookupUuid, err := parseUuidFromParam(c)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	game, responseCode, err := game_client.GetPlatform(lookupUuid)
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

func CreatePlatform(c *gin.Context) {
	var createDto dto.PlatformCreateDto
	if err := c.BindJSON(&createDto); err != nil {
		log.Infof("Failed to parse platform create model: %s", err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	responseBody, responseCode, err := game_client.CreatePlatform(&createDto)
	if err != nil {
		if responseCode == http.StatusNotFound {
			c.AbortWithStatus(http.StatusNotFound)
		} else if responseCode == http.StatusBadRequest {
			c.AbortWithStatus(http.StatusBadRequest)
		} else {
			c.AbortWithStatus(http.StatusInternalServerError)
		}
		return
	}

	c.JSON(http.StatusCreated, responseBody)
}

func UpdatePlatform(c *gin.Context) {
	var updateDto dto.PlatformUpdateDto
	if err := c.BindJSON(&updateDto); err != nil {
		log.Infof("Failed to parse platform update model: %s", err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	id, err := parseUuidFromParam(c)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	responseBody, responseCode, err := game_client.UpdatePlatform(&updateDto, id)
	if err != nil {
		if responseCode == http.StatusNotFound {
			c.AbortWithStatus(http.StatusNotFound)
		} else if responseCode == http.StatusBadRequest {
			c.AbortWithStatus(http.StatusBadRequest)
		} else {
			c.AbortWithStatus(http.StatusInternalServerError)
		}
		return
	}

	c.JSON(http.StatusOK, responseBody)
}

func SetupPlatformRoutes(engine *gin.Engine, basePath string) {
	baseUrl := fmt.Sprintf("%s/api/v0/platforms", basePath)

	engine.GET(baseUrl+"/:id", GetPlatform)
	engine.GET(baseUrl, GetPlatforms)
	engine.POST(baseUrl, CreatePlatform)
	engine.POST(baseUrl+"/:id", UpdatePlatform)
}
