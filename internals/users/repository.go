package users

import "github.com/paper-plane/internals/models"

type Repository interface {
	Register(r *models.RegisterUser) (string, error)
	GetUserResponse(id string) (*models.UserResponse, error)
	UserJWTPayloadCheck(payload *models.UserJWTPayload) (*models.User, error)
	DeleteUser(id string) error
	GetUserPassword(id string) (*string, error)
}
