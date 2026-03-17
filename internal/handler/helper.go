package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/nasrul78/spendly-api/internal/domain"
)

type errorResponse struct {
	Message string `json:"message"`
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, errorResponse{Message: message})
}

func handleError(w http.ResponseWriter, err error) {
	if errors.Is(err, domain.ErrNotFound) {
		writeError(w, http.StatusNotFound, err.Error())
		return
	}

	if errors.Is(err, domain.ErrConflict) {
		writeError(w, http.StatusConflict, err.Error())
		return
	}

	if errors.Is(err, domain.ErrUnauthorized) {
		writeError(w, http.StatusUnauthorized, err.Error())
		return
	}

	writeError(w, http.StatusInternalServerError, err.Error())
}
