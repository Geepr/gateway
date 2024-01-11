package game_client

import (
	"fmt"
	"github.com/Geepr/gateway/clients/game_client/dto"
	"github.com/gofrs/uuid"
	"net/http"
)

// GetRelease returns a release as received from the game microservice.
// The second return value is a response code received from the microservice, or -1 if the request didn't go that far.
func GetRelease(id uuid.UUID) (release *dto.ReleaseDto, responseCode int, err error) {
	path := fmt.Sprintf("v0/releases/%s", id.String())
	return sendAndParseResponse[dto.ReleaseDto](path, http.MethodGet, nil)
}

// GetReleases returns a list of releases as received from the game microservice.
// The second return value is a response code received from the microservice, or -1 if the request didn't go that far.
func GetReleases(page int, size int, gameId uuid.UUID) (releases *dto.ReleaseResponseDto, responseCode int, err error) {
	path := fmt.Sprintf("v0/releases?page=%d&size=%d&gameId=%s", page, size, gameId.String())
	return sendAndParseResponse[dto.ReleaseResponseDto](path, http.MethodGet, nil)
}
