package game_client

import "github.com/gofrs/uuid"

type GameDto struct {
	Id          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description *string   `json:"description"`
	Archived    bool      `json:"archived"`
}
