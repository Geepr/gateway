package game_client

import "github.com/gofrs/uuid"

type GameResponseDto struct {
	Games      []GameDto `json:"games"`
	Page       int       `json:"page"`
	PageSize   int       `json:"pageSize"`
	TotalPages int       `json:"totalPages"`
}

type GameDto struct {
	Id          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description *string   `json:"description"`
	Archived    bool      `json:"archived"`
}

type GameUpdateDto struct {
	Title       string  `json:"title" form:"title" binding:"required"`
	Description *string `json:"description" form:"description"`
	Archived    bool    `json:"archived" form:"archived"`
}

type GameCreateDto struct {
	Title       string  `json:"title" form:"title" binding:"required,max=200"`
	Description *string `json:"description" form:"description" binding:"max=2000"`
}

type GameCreateResponseDto struct {
	Id uuid.UUID `json:"id"`
}
