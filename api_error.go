package main

import (
	"encoding/json"
)

const (
	StatusDatabaseErrorCode = 501
)

type APIError struct {
	ErrorCode             int             `json:"code"`
	ErrorMessage          string          `json:"message"`
	AdditionalInformation json.RawMessage `json:"additionalInformation"`
}
