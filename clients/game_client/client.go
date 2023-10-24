package game_client

import (
	"encoding/json"
	"fmt"
	"github.com/Geepr/gateway/config"
	"github.com/gofrs/uuid"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
)

// GetGame returns a game as received from the game microservice.
// The second return value is a response code received from the microservice, or -1 if the request didn't go that far.
func GetGame(id uuid.UUID) (game *GameDto, responseCode int, err error) {
	path := fmt.Sprintf("v0/games/%s", id.String())
	return sendGetAndParseResponse[GameDto](path)
}

// GetGames returns games list as received from the game microservice.
// The second return value is a response code received from the microservice, or -1 if the request didn't go that far.
func GetGames(page int) (games *[]GameDto, responseCode int, err error) {
	path := fmt.Sprintf("v0/games?page=%d", page)
	return sendGetAndParseResponse[[]GameDto](path)
}

func sendGetAndParseResponse[T any](requestPath string) (parsed *T, responseCode int, err error) {
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/api/%s", config.GameUrl, requestPath), nil)
	if err != nil {
		log.Warnf("Failed to prepare request to: %s due to: %s", requestPath, err.Error())
		return nil, -1, err
	}

	response, err := httpClient.Do(request)
	if err != nil {
		log.Warnf("Failed to send request to: %s due to: %s", requestPath, err.Error())
		return nil, -1, err
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		log.Warnf("Failed to read body from %s due to: %s", requestPath, err.Error())
		return nil, response.StatusCode, err
	}
	var tempParsed T
	err = json.Unmarshal(data, &tempParsed)
	if err != nil {
		log.Warnf("Failed to parse dto from request to %s due to: %s", requestPath, err.Error())
		return nil, response.StatusCode, err
	}
	return &tempParsed, response.StatusCode, nil
}
