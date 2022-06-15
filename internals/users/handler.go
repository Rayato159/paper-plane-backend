package users

import (
	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	Register(c *fiber.Ctx) error
	RemoveUser(c *fiber.Ctx) error
}
