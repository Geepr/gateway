package client_utils

import "errors"

var (
	UnexpectedResponseCode = errors.New("received an unexpected response code")
)
