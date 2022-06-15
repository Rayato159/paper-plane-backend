package auth

import "github.com/paper-plane/internals/models"

type Repository interface {
	FindOneUser(username string) (*models.User, *models.UserCredentials, error)
	CreateJwtToken(userId string, refreshToken string) error
	RefreshToken(refreshToken string) (*models.UserCredentials, error)
}
