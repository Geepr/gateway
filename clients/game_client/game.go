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
