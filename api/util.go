package api

import (
	"encoding/json"
	"net/http"
)

type httpError struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func SendErrorResponse(w http.ResponseWriter, message string, statusCode int) {
	err := &httpError{
		Message: message,
		Status:  http.StatusInternalServerError,
	}

	if statusCode != -1 {
		err.Status = statusCode
	}

	response, _ := json.Marshal(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.Status)
	w.Write(response)
}

func ProcessQueries(params map[string][]string) map[string]string {
	filters := make(map[string]string)
	for key, value := range params {
		if len(value) > 0 {
			filters[key] = value[0]
		}
	}
	return filters
}
