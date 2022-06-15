package usecases

import (
	"errors"
	"log"

	"github.com/paper-plane/internals/accounts"
	"github.com/paper-plane/internals/models"
	"github.com/paper-plane/internals/users"
	"github.com/paper-plane/pkg/utils"

	"github.com/google/uuid"
)

type usersUsecase struct {
	usersRepo    users.Repository
	accountsRepo accounts.Repository
}

func NewUsersUsecase(usersRepo users.Repository, accountsRepo accounts.Repository) users.Usecase {
	return &usersUsecase{
		usersRepo:    usersRepo,
		accountsRepo: accountsRepo,
	}
}

func (u *usersUsecase) Register(r *models.RegisterUser) (*models.UserResponse, error) {
	if r.Password != r.PasswordConfirm {
		return nil, errors.New("error, password doesn't match.")
	}

	hashPassword, err := utils.HashPassword(r.Password)
	if err != nil {
		return nil, err
	}
	r.Password = hashPassword

	// Generate user id
	r.Id = uuid.New().String()

	// Insert User
	id, err := u.usersRepo.Register(r)
	if err != nil {
		return nil, errors.New("error, username have been already using.")
	}

	if err := u.accountsRepo.CreateAccount(id, r.Balance); err != nil {
		log.Println(err.Error())
		if err := u.usersRepo.DeleteUser(id); err != nil {
			log.Println(err.Error())
			return nil, err
		}
		return nil, err
	}

	// Get User Response
	user, err := u.usersRepo.GetUserResponse(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *usersUsecase) RemoveUser(req *models.DeleteUser) error {
	userPassword, err := u.usersRepo.GetUserPassword(req.Id)
	if err != nil {
		return err
	}

	if !utils.ComparePasswordHash(req.Password, *userPassword) {
		return errors.New("error, user not found.")
	}

	if err := u.usersRepo.DeleteUser(req.Id); err != nil {
		return err
	}
	return nil
}

func (u *usersUsecase) UserJWTPayloadCheck(payload *models.UserJWTPayload) error {
	_, err := u.usersRepo.UserJWTPayloadCheck(payload)
	if err != nil {
		return err
	}
	return nil
}
