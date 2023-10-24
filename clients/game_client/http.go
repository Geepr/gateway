package game_client

import (
	"github.com/Geepr/gateway/config"
	"net/http"
)

var httpClient = http.Client{
	Timeout: config.ClientTimeout,
}
