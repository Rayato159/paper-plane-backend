package users

import "github.com/paper-plane/internals/models"

type Usecase interface {
	Register(r *models.RegisterUser) (*models.UserResponse, error)
	RemoveUser(req *models.DeleteUser) error
	UserJWTPayloadCheck(payload *models.UserJWTPayload) error
}
