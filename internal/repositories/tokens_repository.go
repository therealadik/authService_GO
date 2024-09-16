package repositories

import (
	"github.com/google/uuid"
	"github.com/therealadik/auth-service/internal/database"
	"github.com/therealadik/auth-service/internal/models"
)

func GetRefreshToken(tokenID uuid.UUID) (*models.RefreshTokenData, error) {
	query := `
        SELECT token_hash, ip_address 
        FROM refresh_tokens 
        WHERE token_id = $1;
    `
	var tokenData models.RefreshTokenData

	err := database.DB.QueryRow(query, tokenID).Scan(&tokenData.TokenHash, &tokenData.IPAddress)
	if err != nil {
		return nil, err
	}

	return &tokenData, nil
}

func SaveRefreshToken(tokenID uuid.UUID, tokenHash, ipAddress string) error {
	query := `
	INSERT INTO refresh_tokens (token_id, token_hash, ip_address) 
	VALUES ($1, $2, $3)
	ON CONFLICT (ip_address) 
	DO UPDATE SET token_hash = EXCLUDED.token_hash, token_id = EXCLUDED.token_id;
`

	_, err := database.DB.Exec(query, tokenID, tokenHash, ipAddress)
	return err
}
