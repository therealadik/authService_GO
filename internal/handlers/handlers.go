package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/therealadik/auth-service/internal/models"
	"github.com/therealadik/auth-service/internal/services"
)

func AuthHandler(w http.ResponseWriter, r *http.Request) {
	var req models.TokenRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response, err := services.AuthUser(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func RefreshTokens(w http.ResponseWriter, r *http.Request) {
	var req models.RefreshRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	accessToken, refreshToken, err := services.RefreshTokens(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	response := models.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
