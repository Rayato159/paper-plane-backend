package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/paper-plane/internals/models"
)

func JwtClaimsAccessToken(p *models.Payload) (*string, error) {
	mySigningKey := []byte(os.Getenv("JWT_SECRET_KEY"))

	claims := models.Payload{
		p.UserId,
		p.Username,
		p.AccountId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		return nil, err
	}
	return &ss, nil
}

func JwtClaimsRefreshToken(p *models.Payload) (*string, error) {
	mySigningKey := []byte(os.Getenv("JWT_SECRET_KEY"))

	claims := models.Payload{
		p.UserId,
		p.Username,
		p.AccountId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		return nil, err
	}
	return &ss, nil
}

func ExtractPayload(JWTtoken string) (*models.Payload, error) {
	return nil, nil
}
