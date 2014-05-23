package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type APIResponse struct {
	Data         json.RawMessage `json:"data,omitempty"`
	Error        *APIError       `json:"error,omitempty"`
	ResponseCode int             `json:"-"`
}

func (a *APIResponse) setResponseCode(code int) {
	a.ResponseCode = code
}

func (a *APIResponse) send(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	if a.ResponseCode != 0 {
		w.WriteHeader(a.ResponseCode)
	}
	w.Write(a.getJSON())
}

func (a *APIResponse) getJSON() []byte {
	b, err := json.Marshal(a)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return b
}
