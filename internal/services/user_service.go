package services

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/therealadik/auth-service/internal/repositories"
)

func SendEmailMessage(userID uuid.UUID, message string) error {
	user, err := repositories.GetUserByID(userID)
	if err != nil {
		return err
	}

	fmt.Printf("Email sent to: %s. Message: %s\n", user.Email, message)
	return nil
}
