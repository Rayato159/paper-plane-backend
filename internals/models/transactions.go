package models

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	Id        uuid.UUID `db:"id" json:"id"`
	Income    float64   `db:"income" json:"income"`
	Expense   float64   `db:"expense" json:"expense"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	AccountId string    `db:"account_id" json:"account_id"`
}

type CreateReqTransaction struct {
	Id        string  `db:"id" json:"id"`
	Income    float64 `db:"income" json:"income" from:"income"`
	Expense   float64 `db:"expense" json:"expense" from:"expense"`
	AccountId string  `db:"account_id" json:"account_id"`
	Balance   float64 `db:"balance" json:"balance"`
}

type EditReqTransaction struct {
	Id      string  `db:"id" json:"id"`
	Income  float64 `db:"income" json:"income" from:"income"`
	Expense float64 `db:"expense" json:"expense" from:"expense"`
	Balance float64 `db:"balance" json:"balance"`
}

type ReqTransactionQuery struct {
	StartDate   string             `query:"start"`
	EndDate     string             `query:"end"`
	Sort        TransactionSort    `query:"sort"`
	OrderBy     TransactionOrderBy `query:"order_by"`
	Page        int                `query:"page"`
	ItemPerPage int                `query:"item_per_page"`
}

type TransactionPagination struct {
	Page        int            `json:"page"`
	ItemPerPage int            `json:"item_per_page"`
	TotalItem   int            `json:"total_item"`
	Data        []*Transaction `json:"data"`
}

type TransactionRemoveReq struct {
	Id       string `db:"id" json:"id"`
	UserId   string `db:"user_id" json:"user_id" form:"user_id"`
	Password string `db:"password" json:"password" form:"password"`
}

// Sort Enum
type TransactionSort string

const (
	ASC  TransactionSort = "ASC"
	DESC TransactionSort = "DESC"
)

// Order By Enum
type TransactionOrderBy string

const (
	DATE    TransactionOrderBy = "created_date"
	UPDATED TransactionOrderBy = "updated_date"
	INCOME  TransactionOrderBy = "income"
	EXPENSE TransactionOrderBy = "expense"
)
