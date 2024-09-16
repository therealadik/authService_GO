package repositories

import (
	"github.com/google/uuid"
	"github.com/therealadik/auth-service/internal/database"
	"github.com/therealadik/auth-service/internal/models"
)

func GetUserByID(userID uuid.UUID) (*models.User, error) {
	query := `
	SELECT id, email FROM users WHERE id = $1;
`
	var user models.User

	err := database.DB.QueryRow(query, userID).Scan(&user.ID, &user.Email)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func CheckUserExists(userID uuid.UUID) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE id=$1)"
	err := database.DB.QueryRow(query, userID).Scan(&exists)
	return exists, err
}
