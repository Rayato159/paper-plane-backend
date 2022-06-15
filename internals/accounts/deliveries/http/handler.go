package http

import (
	"errors"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/paper-plane/internals/accounts"
)

type accountsHandler struct {
	accountsUC accounts.Usecase
}

func NewAccountsHandler(accountsUC accounts.Usecase) accounts.Handler {
	return &accountsHandler{accountsUC: accountsUC}
}

func (h *accountsHandler) GetAccountInfo(c *fiber.Ctx) error {
	accountId := c.Params("id")
	userId := c.Locals("id").(string)

	res, err := h.accountsUC.GetAccountInfo(accountId, userId)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":      "Internal Server Error",
			"status_code": fiber.StatusInternalServerError,
			"message":     err.Error(),
			"result":      nil,
		})
	}

	if res.Id == "" {
		log.Println(errors.New("error, account not found.").Error())
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":      "No content",
			"status_code": fiber.StatusNoContent,
			"message":     "error, account not found.",
			"result":      nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":      "OK",
		"status_code": fiber.StatusOK,
		"message":     nil,
		"result":      res,
	})
}
