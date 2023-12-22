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
