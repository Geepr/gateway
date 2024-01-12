package dto

import (
	"github.com/gofrs/uuid"
	"time"
)

type ReleaseResponseDto struct {
	Releases   []ReleaseDto `json:"releases"`
	Page       int          `json:"page"`
	PageSize   int          `json:"pageSize"`
	TotalPages int          `json:"totalPages"`
}

type ReleaseDto struct {
	Id                 uuid.UUID  `json:"id"`
	GameId             uuid.UUID  `json:"gameId"`
	TitleOverride      *string    `json:"titleOverride"`
	Description        *string    `json:"description"`
	ReleaseDate        *time.Time `json:"releaseDate"`
	ReleaseDateUnknown bool       `json:"releaseDateUnknown"`
}

type ReleaseUpdateDto struct {
	TitleOverride      *string    `json:"titleOverride"`
	Description        *string    `json:"description"`
	ReleaseDate        *time.Time `json:"releaseDate"`
	ReleaseDateUnknown bool       `json:"releaseDateUnknown"`
}

type ReleaseCreateDto struct {
	GameId             uuid.UUID  `json:"gameId"`
	TitleOverride      *string    `json:"titleOverride"`
	Description        *string    `json:"description"`
	ReleaseDate        *time.Time `json:"releaseDate"`
	ReleaseDateUnknown bool       `json:"releaseDateUnknown"`
}

type ReleaseCreateResponseDto struct {
	Id uuid.UUID `json:"id"`
}
