package services

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"log"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/therealadik/auth-service/internal/config"
	"github.com/therealadik/auth-service/internal/models"
	"github.com/therealadik/auth-service/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

func AuthUser(req models.TokenRequest) (models.TokenResponse, error) {

	userExist, err := repositories.CheckUserExists(req.UserID)
	if err != nil {
		return models.TokenResponse{}, err
	}

	if !userExist {
		return models.TokenResponse{}, errors.New("user not found")
	}

	accessToken, refreshToken, err := generateAndSaveTokenPair(req.UserID, req.ClientIP)
	if err != nil {
		return models.TokenResponse{}, err
	}

	response := models.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return response, nil
}

func RefreshTokens(req models.RefreshRequest) (string, string, error) {
	accessTokenClaims, err := parseAccessToken(req.AccessToken)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := repositories.GetRefreshToken(accessTokenClaims.TokenID)
	if err != nil {
		return "", "", err
	}

	err = isValidRefreshToken(refreshToken.TokenHash, req.RefreshToken)
	if err != nil {
		return "", "", err
	}

	if refreshToken.IPAddress != accessTokenClaims.IPAddress {
		err := SendEmailMessage(accessTokenClaims.UserID, "ip address changed")
		if err != nil {
			log.Print(err)
		}
		return "", "", errors.New("ip changed")
	}

	newAccessToken, newRefreshToken, err := generateAndSaveTokenPair(accessTokenClaims.UserID, accessTokenClaims.IPAddress)

	return newAccessToken, newRefreshToken, err
}

func generateRefreshToken() (string, string, error) {
	token := make([]byte, 32)
	_, err := rand.Read(token)
	if err != nil {
		return "", "", err
	}

	refreshToken := base64.StdEncoding.EncodeToString(token)
	hashedToken, err := bcrypt.GenerateFromPassword([]byte(refreshToken), bcrypt.DefaultCost)

	return refreshToken, string(hashedToken), err
}

func generateAccessToken(userID uuid.UUID, ipAddress string, tokenID uuid.UUID) (string, error) {
	claims := models.Claims{
		UserID:    userID,
		IPAddress: ipAddress,
		TokenID:   tokenID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	return token.SignedString([]byte(config.JwtSecret))
}

func parseAccessToken(accessToken string) (*models.Claims, error) {
	parsedToken, err := jwt.ParseWithClaims(accessToken, &models.Claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.JwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := parsedToken.Claims.(*models.Claims)
	if !ok || !parsedToken.Valid {
		return nil, fmt.Errorf("invalid token or claims")
	}

	return claims, nil
}

func generateAndSaveTokenPair(userID uuid.UUID, ipAddress string) (string, string, error) {
	tokenID := uuid.New()

	accessToken, err := generateAccessToken(userID, ipAddress, tokenID)
	if err != nil {
		return "", "", err
	}

	refreshToken, hashedToken, err := generateRefreshToken()
	if err != nil {
		return "", "", err
	}

	err = repositories.SaveRefreshToken(tokenID, hashedToken, ipAddress)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func isValidRefreshToken(tokenHash, refreshToken string) error {
	err := bcrypt.CompareHashAndPassword([]byte(tokenHash), []byte(refreshToken))
	if err != nil {
		return errors.New("refreshToken not valid")
	}

	return nil
}
