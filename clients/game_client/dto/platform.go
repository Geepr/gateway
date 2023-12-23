package dto

import "github.com/gofrs/uuid"

type PlatformDto struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	ShortName string    `json:"shortName"`
}

type PlatformResponseDto struct {
	Platforms  []*PlatformDto `json:"platforms"`
	Page       int            `json:"page"`
	PageSize   int            `json:"pageSize"`
	TotalPages int            `json:"totalPages"`
}

type PlatformUpdateDto struct {
	Name      string `json:"name" form:"name" binding:"required,max=200"`
	ShortName string `json:"shortName" form:"shortName" binding:"required,max=10"`
}

type PlatformCreateDto struct {
	Name      string `json:"name" form:"name" binding:"required,max=200"`
	ShortName string `json:"shortName" form:"shortName" binding:"required,max=10"`
}

type PlatformCreateResponseDto struct {
	Id uuid.UUID `json:"id"`
}
