package http

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/paper-plane/internals/models"
	"github.com/paper-plane/internals/transactions"
)

type transactionsHandler struct {
	transactionsUC transactions.Usecase
}

func NewTransactionsHandler(transactionsUC transactions.Usecase) transactions.Handler {
	return &transactionsHandler{transactionsUC: transactionsUC}
}

func (h *transactionsHandler) AddTransaction(c *fiber.Ctx) error {
	accountId := c.Params("account_id")
	req := new(models.CreateReqTransaction)
	userId := c.Locals("id").(string)

	if err := c.BodyParser(req); err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":      "Bad Request",
			"status_code": fiber.StatusBadRequest,
			"message":     err.Error(),
			"result":      nil,
		})
	}
	req.AccountId = accountId

	res, err := h.transactionsUC.AddTransaction(req, userId)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":      "Internal Server Error",
			"status_code": fiber.StatusInternalServerError,
			"message":     "error, can't insert a transaction.",
			"result":      nil,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":      "Created",
		"status_code": fiber.StatusCreated,
		"message":     nil,
		"result":      res,
	})
}

func (h *transactionsHandler) GetTransactionLists(c *fiber.Ctx) error {
	queryReq := new(models.ReqTransactionQuery)
	accountId := c.Params("account_id")
	userId := c.Locals("id").(string)

	if err := c.QueryParser(queryReq); err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":      "Internal Server Error",
			"status_code": fiber.StatusInternalServerError,
			"message":     err.Error(),
			"result":      nil,
		})
	}

	res, err := h.transactionsUC.GetTransactionLists(accountId, queryReq, userId)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":      "Internal Server Error",
			"status_code": fiber.StatusInternalServerError,
			"message":     err.Error(),
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

func (h *transactionsHandler) GetTransactionById(c *fiber.Ctx) error {
	transactionId := c.Params("id")
	userId := c.Locals("id").(string)

	res, err := h.transactionsUC.GetTransactionById(transactionId, userId)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":      "Internal Server Error",
			"status_code": fiber.StatusInternalServerError,
			"message":     err.Error(),
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

func (h *transactionsHandler) EditTransaction(c *fiber.Ctx) error {
	req := new(models.EditReqTransaction)
	transactionId := c.Params("id")
	userId := c.Locals("id").(string)

	if err := c.BodyParser(req); err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":      "Bad Request",
			"status_code": fiber.StatusBadRequest,
			"message":     err.Error(),
			"result":      nil,
		})
	}

	res, err := h.transactionsUC.EditTransaction(transactionId, req, userId)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":      "Internal Server Error",
			"status_code": fiber.StatusInternalServerError,
			"message":     "error, can't edit transaction.",
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

func (h *transactionsHandler) RemoveTransaction(c *fiber.Ctx) error {
	transactionId := c.Params("id")
	userId := c.Locals("id").(string)
	req := new(models.TransactionRemoveReq)
	if err := c.BodyParser(req); err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":      "Bad Request",
			"status_code": fiber.StatusBadRequest,
			"message":     err.Error(),
			"result":      nil,
		})
	}
	req.Id = transactionId
	req.UserId = userId

	if err := h.transactionsUC.RemoveTransaction(req, userId); err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":      "Internal Server Error",
			"status_code": fiber.StatusInternalServerError,
			"message":     "error, hey hey stop right there!",
			"result":      nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":      "No Content",
		"status_code": fiber.StatusNoContent,
		"message":     nil,
		"result":      nil,
	})
}
