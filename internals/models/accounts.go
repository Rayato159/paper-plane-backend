package models

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	Id        uuid.UUID `db:"id" json:"id"`
	Balance   float64   `db:"balance" json:"balance"`
	UserId    uuid.UUID `db:"user_id" json:"user_id"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type AccountInfo struct {
	Id        string       `db:"id" json:"id"`
	Balance   float64      `db:"balance" json:"balance"`
	User      UserResponse `json:"user"`
	CreatedAt time.Time    `db:"created_at" json:"created_at"`
	UpdatedAt time.Time    `db:"updated_at" json:"updated_at"`
}

type Balance struct {
	Balance float64 `db:"balance" json:"balance"`
}
