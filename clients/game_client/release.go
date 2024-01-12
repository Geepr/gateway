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

func UpdateRelease(id uuid.UUID, release *dto.ReleaseUpdateDto) (responseCode int, err error) {
	path := fmt.Sprintf("v0/releases/%s", id.String())
	releaseJson, err := json.Marshal(release)
	if err != nil {
		return -1, err
	}
	response, err := send(path, http.MethodPut, bytes.NewBuffer(releaseJson))
	if err != nil {
		log.Warnf("Failed to put updated release: %s", err.Error())
		return -1, err
	}
	_ = response.Body.Close()
	return response.StatusCode, nil
}

func CreateRelease(release *dto.ReleaseCreateDto) (response *dto.ReleaseCreateResponseDto, responseCode int, err error) {
	path := fmt.Sprintf("v0/releases/")
	releaseJson, err := json.Marshal(release)
	if err != nil {
		return nil, -1, err
	}
	return sendAndParseResponse[dto.ReleaseCreateResponseDto](path, http.MethodPost, bytes.NewBuffer(releaseJson))
}

func DeleteRelease(id uuid.UUID) (responseCode int, err error) {
	path := fmt.Sprintf("v0/releases/%s", id.String())
	response, err := send(path, http.MethodDelete, nil)
	if err != nil {
		log.Warnf("Failed to delete release: %s", err.Error())
		return -1, err
	}
	_ = response.Body.Close()
	return response.StatusCode, nil
}
