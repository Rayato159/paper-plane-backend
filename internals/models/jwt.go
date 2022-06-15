package models

import "github.com/golang-jwt/jwt/v4"

type Payload struct {
	UserId    string `json:"id"`
	Username  string `json:"username"`
	AccountId string `json:"account_id"`
	jwt.StandardClaims
}

type UserJWTPayload struct {
	UserId   string  `db:"id" json:"id"`
	Username string  `db:"username" json:"username"`
	Expires  float64 `db:"exp" json:"exp"`
}
