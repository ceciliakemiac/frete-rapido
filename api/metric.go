package api

import (
	"encoding/json"
	"net/http"
)

func (s *Server) GetMetrics(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	filters := ProcessQueries(params)

	metrics, err := s.db.GetMetrics(filters)
	if err != nil {
		SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(metrics)
	if err != nil {
		SendErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
