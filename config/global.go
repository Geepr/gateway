package config

import "time"

var (
	GameUrl       string = "http://localhost:5500" //todo: read from env
	ClientTimeout        = time.Second * 30
)
