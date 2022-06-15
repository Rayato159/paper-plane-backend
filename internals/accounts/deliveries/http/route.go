package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/paper-plane/internals/accounts"
	"github.com/paper-plane/internals/users"
	"github.com/paper-plane/pkg/middlewares"
)

func MapAccountsRoute(r fiber.Router, h accounts.Handler, u users.Usecase) {
	r.Get("/:id", middlewares.JwtAuthentication(u), h.GetAccountInfo)
}
