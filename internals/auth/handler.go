package auth

import "github.com/gofiber/fiber/v2"

type Handler interface {
	Login(c *fiber.Ctx) error
	RefreshToken(c *fiber.Ctx) error
}
