package auth

import "github.com/paper-plane/internals/models"

type Usecase interface {
	Login(m *models.Credentials) (*models.UserCredentials, error)
	RefreshToken(refreshToken string) (*models.UserCredentials, error)
}
