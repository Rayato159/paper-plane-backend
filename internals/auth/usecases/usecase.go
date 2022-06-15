package usecases

import (
	"errors"

	"github.com/paper-plane/internals/auth"
	"github.com/paper-plane/internals/models"
	"github.com/paper-plane/pkg/utils"
)

type authUsecase struct {
	authRepo auth.Repository
}

func NewAuthUsecase(authRepo auth.Repository) auth.Usecase {
	return &authUsecase{authRepo: authRepo}
}

func (u *authUsecase) Login(m *models.Credentials) (*models.UserCredentials, error) {
	user, credentials, err := u.authRepo.FindOneUser(m.Username)
	if err != nil {
		return nil, errors.New("error, user not found.")
	}

	if user == nil || !utils.ComparePasswordHash(m.Password, user.Password) {
		return nil, errors.New("error, username or password is incorrect.")
	}

	payload := new(models.Payload)
	payload.UserId = user.Id.String()
	payload.Username = user.Username
	payload.AccountId = *user.AccountId

	accessTokenChan := make(chan *string)
	refreshTokenChan := make(chan *string)
	accessTokenErrChan := make(chan error)
	refreshTokenErrChan := make(chan error)

	go func(c chan *string, e chan error) {
		accessToken, err := utils.JwtClaimsAccessToken(payload)
		accessTokenErrChan <- err
		accessTokenChan <- accessToken
		close(accessTokenErrChan)
		close(accessTokenChan)
	}(accessTokenChan, accessTokenErrChan)

	go func(c chan *string, e chan error) {
		refreshToken, err := utils.JwtClaimsRefreshToken(payload)
		refreshTokenErrChan <- err
		refreshTokenChan <- refreshToken
		close(refreshTokenErrChan)
		close(refreshTokenChan)
	}(refreshTokenChan, refreshTokenErrChan)

	accessTokenErr := <-accessTokenErrChan
	refreshTokenErr := <-refreshTokenErrChan

	if accessTokenErr != nil {
		return nil, accessTokenErr
	}

	if refreshTokenErr != nil {
		return nil, refreshTokenErr
	}

	credentials.AccessToken = <-accessTokenChan
	credentials.RefreshToken = *<-refreshTokenChan

	if err := u.authRepo.CreateJwtToken(credentials.Id, credentials.RefreshToken); err != nil {
		return nil, err
	}
	return credentials, nil
}

func (u *authUsecase) RefreshToken(refreshToken string) (*models.UserCredentials, error) {
	credentials, err := u.authRepo.RefreshToken(refreshToken)
	if err != nil {
		return nil, errors.New("error, user not found.")
	}

	accessTokenChan := make(chan *string)
	accessTokenErrChan := make(chan error)

	payload := new(models.Payload)
	payload.UserId = credentials.Id
	payload.Username = credentials.Username
	payload.AccountId = credentials.AccountId

	go func(c chan *string, e chan error) {
		accessToken, err := utils.JwtClaimsAccessToken(payload)
		accessTokenErrChan <- err
		accessTokenChan <- accessToken
		close(accessTokenErrChan)
		close(accessTokenChan)
	}(accessTokenChan, accessTokenErrChan)

	accessTokenErr := <-accessTokenErrChan
	if accessTokenErr != nil {
		return nil, accessTokenErr
	}
	credentials.AccessToken = <-accessTokenChan

	return credentials, nil
}
