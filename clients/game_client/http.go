package game_client

import (
	"encoding/json"
	"fmt"
	"github.com/Geepr/gateway/clients/client_utils"
	"github.com/Geepr/gateway/config"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
)

var httpClient = http.Client{
	Timeout: config.ClientTimeout,
}

func send(requestPath string, method string, body io.Reader) (response *http.Response, err error) {
	request, err := http.NewRequest(method, fmt.Sprintf("%s/api/%s", config.GameUrl, requestPath), body)
	if err != nil {
		log.Warnf("Failed to prepare request to: %s due to: %s", requestPath, err.Error())
		return nil, err
	}

	response, err = httpClient.Do(request)
	if err != nil {
		log.Warnf("Failed to send request to: %s due to: %s", requestPath, err.Error())
		return nil, err
	}
	return response, nil
}

func sendAndParseResponse[T any](requestPath string, method string, body io.Reader) (parsed *T, responseCode int, err error) {
	response, err := send(requestPath, method, body)
	if err != nil {
		return nil, -1, err
	}
	if response.StatusCode < 200 || response.StatusCode > 299 {
		return nil, response.StatusCode, client_utils.UnexpectedResponseCode
	}
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
