package lib

import (
	"encoding/json"
	"log"
	"net/http"
)

// ErrorResponse represents the structure of an error response
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message,omitempty"`
	Code    int    `json:"code"`
}

// ResponseWithJSON sends a JSON response with the given status code and payload
func ResponseWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON response: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, err = w.Write(response)
	if err != nil {
		log.Printf("Error writing response: %v", err)
	}
}

// ResponseWithError sends a JSON error response with the given status code and error message
func ResponseWithError(w http.ResponseWriter, statusCode int, message string) {
	errorResponse := ErrorResponse{
		Error:   http.StatusText(statusCode),
		Message: message,
		Code:    statusCode,
	}
	ResponseWithJSON(w, statusCode, errorResponse)
}
