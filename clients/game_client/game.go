package game_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Geepr/gateway/clients/game_client/dto"
	"github.com/gofrs/uuid"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// GetGame returns a game as received from the game microservice.
// The second return value is a response code received from the microservice, or -1 if the request didn't go that far.
func GetGame(id uuid.UUID) (game *dto.GameDto, responseCode int, err error) {
	path := fmt.Sprintf("v0/games/%s", id.String())
	return sendAndParseResponse[dto.GameDto](path, http.MethodGet, nil)
}

// GetGames returns games list as received from the game microservice.
// The second return value is a response code received from the microservice, or -1 if the request didn't go that far.
func GetGames(page int, size int, title string) (games *dto.GameResponseDto, responseCode int, err error) {
	path := fmt.Sprintf("v0/games?page=%d&size=%d&title=%s", page, size, title)
	return sendAndParseResponse[dto.GameResponseDto](path, http.MethodGet, nil)
}

func UpdateGame(id uuid.UUID, game *dto.GameUpdateDto) (responseCode int, err error) {
	path := fmt.Sprintf("v0/games/%s", id.String())
	gameJson, err := json.Marshal(game)
	if err != nil {
		return -1, err
	}
	response, err := send(path, http.MethodPut, bytes.NewBuffer(gameJson))
	if err != nil {
		log.Warnf("Failed to put updated game: %s", err.Error())
		return -1, err
	}
	_ = response.Body.Close()
	return response.StatusCode, nil
}

func CreateGame(game *dto.GameCreateDto) (response *dto.GameCreateResponseDto, responseCode int, err error) {
	path := fmt.Sprintf("v0/games/")
	gameJson, err := json.Marshal(game)
	if err != nil {
		return nil, -1, err
	}
	return sendAndParseResponse[dto.GameCreateResponseDto](path, http.MethodPost, bytes.NewBuffer(gameJson))
}

func DeleteGame(id uuid.UUID) (responseCode int, err error) {
	path := fmt.Sprintf("v0/games/%s", id.String())
	response, err := send(path, http.MethodDelete, nil)
	if err != nil {
		log.Warnf("Failed to delete game: %s", err.Error())
		return -1, err
	}
	_ = response.Body.Close()
	return response.StatusCode, nil
}
