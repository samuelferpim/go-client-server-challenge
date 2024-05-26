package model

import "encoding/json"

type ErrorResponse struct {
	Error string `json:"error"`
}

func ErrorResponseJson(err error) []byte {
	response := ErrorResponse{Error: err.Error()}
	data, _ := json.Marshal(response)
	return data
}
