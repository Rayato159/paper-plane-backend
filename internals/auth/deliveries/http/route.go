package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/paper-plane/internals/auth"
)

func MapAuthRoute(r fiber.Router, h auth.Handler) {
	r.Post("/", h.RefreshToken)
	r.Post("/login", h.Login)
}
