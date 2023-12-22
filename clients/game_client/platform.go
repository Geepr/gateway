package game_client

import (
	"fmt"
	"github.com/Geepr/gateway/clients/game_client/dto"
	"net/http"
)

// GetPlatforms returns platforms list as received from the game microservice.
// The second return value is a response code received from the microservice, or -1 if the request didn't go that far.
func GetPlatforms(page int, size int, name string) (platforms *dto.PlatformResponseDto, responseCode int, err error) {
	path := fmt.Sprintf("v0/platforms?page=%d&size=%d&name=%s", page, size, name)
	return sendAndParseResponse[dto.PlatformResponseDto](path, http.MethodGet, nil)
}
