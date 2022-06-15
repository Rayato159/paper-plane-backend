package models

type Credentials struct {
	Username string `db:"username" json:"username" form:"username"`
	Password string `db:"password" json:"password" form:"password"`
}

type RefreshToken struct {
	RefreshToken string `db:"refresh_token" json:"refresh_token" form:"refresh_token"`
}
