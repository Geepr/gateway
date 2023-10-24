package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	log "github.com/sirupsen/logrus"
)

type id struct {
	Id string `form:"id" uri:"id" binding:"required,uuid"`
}

func parseUuidFromParam(c *gin.Context) (uuid.UUID, error) {
	var id id
	err := c.BindUri(&id)
	if err != nil {
		log.Infof("Failed to parse uuid: %s", err.Error())
		return uuid.Nil, err
	}
	return uuid.FromString(id.Id)
}
