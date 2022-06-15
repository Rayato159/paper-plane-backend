package transactions

import "github.com/gofiber/fiber/v2"

type Handler interface {
	AddTransaction(c *fiber.Ctx) error
	GetTransactionLists(c *fiber.Ctx) error
	GetTransactionById(c *fiber.Ctx) error
	EditTransaction(c *fiber.Ctx) error
	RemoveTransaction(c *fiber.Ctx) error
}
