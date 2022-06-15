package accounts

import "github.com/gofiber/fiber/v2"

type Handler interface {
	GetAccountInfo(c *fiber.Ctx) error
}
