package controllers

import (
	"fmt"
	"github.com/Geepr/gateway/clients/game_client"
	"github.com/Geepr/gateway/clients/game_client/dto"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func GetReleases(c *gin.Context) {
	var query struct {
		PageIndex int    `form:"page" binding:"required"`
		PageSize  int    `form:"pageSize" binding:"required"`
		GameId    string `form:"gameId" binding:"required"`
	}
	var gameIdUuid uuid.UUID
	var err error
	if err, gameIdUuid = c.BindQuery(&query), uuid.FromStringOrNil(query.GameId); err != nil || gameIdUuid == uuid.Nil {
		log.Infof("Failed to bind release query: %s", err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	releases, _, err := game_client.GetReleases(query.PageIndex, query.PageSize, gameIdUuid)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, releases)
}

func GetRelease(c *gin.Context) {
	lookupUuid, err := parseUuidFromParam(c)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	release, responseCode, err := game_client.GetRelease(lookupUuid)
	if err != nil {
		if responseCode == http.StatusNotFound {
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			c.AbortWithStatus(http.StatusInternalServerError)
		}
		return
	}

	c.JSON(http.StatusOK, release)
}

func UpdateRelease(c *gin.Context) {
	var updateDto dto.ReleaseUpdateDto
	if err := c.BindJSON(&updateDto); err != nil {
		log.Infof("Failed to parse release update model: %s", err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	id, err := parseUuidFromParam(c)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	responseCode, err := game_client.UpdateRelease(id, &updateDto)
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

func CreateRelease(c *gin.Context) {
	var createDto dto.ReleaseCreateDto
	if err := c.BindJSON(&createDto); err != nil {
		log.Infof("Failed to parse release create model: %s", err.Error())
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	responseBody, responseCode, err := game_client.CreateRelease(&createDto)
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

func DeleteRelease(c *gin.Context) {
	lookupUuid, err := parseUuidFromParam(c)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	responseCode, err := game_client.DeleteRelease(lookupUuid)
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

func SetupReleaseRoutes(engine *gin.Engine, basePath string) {
	baseUrl := fmt.Sprintf("%s/api/v0/releases", basePath)

	engine.GET(baseUrl, GetReleases)
	engine.GET(baseUrl+"/:id", GetRelease)
	engine.POST(baseUrl+"/:id", UpdateRelease)
	engine.POST(baseUrl, CreateRelease)
	engine.DELETE(baseUrl+"/:id", DeleteRelease)
}
