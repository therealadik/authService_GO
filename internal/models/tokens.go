package models

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Claims struct {
	UserID    uuid.UUID `json:"user_id"`
	IPAddress string    `json:"ip_address"`
	TokenID   uuid.UUID `json:"token_id"`
	jwt.RegisteredClaims
}

type TokenRequest struct {
	UserID   uuid.UUID `json:"user_id"`
	ClientIP string    `json:"client_ip"`
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenData struct {
	TokenHash string
	IPAddress string
}

type RefreshRequest struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
