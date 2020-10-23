package v1

import (
	"net/http"
)

// Const ...
const (
	HTTPStatusOK                  = http.StatusOK
	HTTPStatusBadRequest          = http.StatusBadRequest
	HTTPStatusNotFound            = http.StatusNotFound
	HTTPStatusInternalServerError = http.StatusInternalServerError
)

// GetResponse ...
type GetResponse struct {
	Data      Anagrams `json:"data"`
	Error     bool     `json:"error"`
	ErrorText string   `json:"errorText"`
}

// Anagrams ...
type Anagrams struct {
	Anagram []string `json:"anagram"`
}

// ErrMap ...
var (
	ErrMap = map[string]string{
		"1": "service error",
		"2": "authorization error",
		"3": "wrong query parameter",
		"4": "missing required parameter",
		"5": "Данные для отчета еще не сформированы",
	}
)
