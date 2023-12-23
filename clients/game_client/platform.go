package game_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Geepr/gateway/clients/game_client/dto"
	"github.com/gofrs/uuid"
	"net/http"
)

// GetPlatform returns a platform as received from the game microservice.
// The second return value is a response code received from the microservice, or -1 if the request didn't go that far.
func GetPlatform(id uuid.UUID) (platform *dto.PlatformDto, responseCode int, err error) {
	path := fmt.Sprintf("v0/platforms/%s", id.String())
	return sendAndParseResponse[dto.PlatformDto](path, http.MethodGet, nil)
}

// GetPlatforms returns platforms list as received from the game microservice.
// The second return value is a response code received from the microservice, or -1 if the request didn't go that far.
func GetPlatforms(page int, size int, name string) (platforms *dto.PlatformResponseDto, responseCode int, err error) {
	path := fmt.Sprintf("v0/platforms?page=%d&size=%d&name=%s", page, size, name)
	return sendAndParseResponse[dto.PlatformResponseDto](path, http.MethodGet, nil)
}

func CreatePlatform(platform *dto.PlatformCreateDto) (response *dto.PlatformCreateResponseDto, responseCode int, err error) {
	path := fmt.Sprintf("v0/platforms/")
	platformJson, err := json.Marshal(platform)
	if err != nil {
		return nil, -1, err
	}
	return sendAndParseResponse[dto.PlatformCreateResponseDto](path, http.MethodPost, bytes.NewBuffer(platformJson))
}
