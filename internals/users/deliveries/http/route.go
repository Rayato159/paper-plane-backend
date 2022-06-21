package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/paper-plane/internals/users"
	"github.com/paper-plane/pkg/middlewares"
)

func MapUsersRoute(r fiber.Router, h users.Handler, u users.Usecase) {
	r.Post("/register", h.Register)
	r.Delete("/", middlewares.JwtAuthentication(u), h.RemoveUser)
}
