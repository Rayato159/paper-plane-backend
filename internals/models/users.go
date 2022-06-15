package models

import "github.com/google/uuid"

type User struct {
	Id           uuid.UUID `db:"id" json:"id"`
	Username     string    `db:"username" json:"username"`
	Password     string    `db:"password" json:"password"`
	AccountId    *string   `db:"account_id" json:"account_id"`
	RefreshToken string    `db:"refresh_token" json:"refresh_token"`
}

type RegisterUser struct {
	Id              string  `db:"id" json:"id"`
	Username        string  `db:"username" json:"username" form:"username"`
	Password        string  `db:"password" json:"password" form:"password"`
	PasswordConfirm string  `json:"password_confirm" form:"password_confirm"`
	Balance         float64 `db:"balance" json:"balance"`
}

type UserResponse struct {
	Id       string `db:"id" json:"id"`
	Username string `db:"username" json:"username"`
}

type UserCredentials struct {
	Id           string  `db:"id" json:"id"`
	Username     string  `db:"username" json:"username"`
	AccountId    string  `db:"account_id" json:"account_id"`
	AccessToken  *string `json:"access_token"`
	RefreshToken string  `db:"refresh_token" json:"refresh_token"`
}

type DeleteUser struct {
	Id       string `db:"id" json:"id"`
	Password string `db:"password" json:"password" form:"password"`
}
