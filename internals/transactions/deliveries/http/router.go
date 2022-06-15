package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/paper-plane/internals/transactions"
	"github.com/paper-plane/internals/users"
	"github.com/paper-plane/pkg/middlewares"
)

func MapTransactionsRoute(r fiber.Router, h transactions.Handler, u users.Usecase) {
	r.Post(":account_id", middlewares.JwtAuthentication(u), h.AddTransaction)
	r.Get(":account_id", middlewares.JwtAuthentication(u), h.GetTransactionLists)
	r.Get(":id/info", middlewares.JwtAuthentication(u), h.GetTransactionById)
	r.Put(":id", middlewares.JwtAuthentication(u), h.EditTransaction)
	r.Delete(":id", middlewares.JwtAuthentication(u), h.RemoveTransaction)
}
