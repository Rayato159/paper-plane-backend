package http

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/paper-plane/internals/models"
	"github.com/paper-plane/internals/users"
)

type usersHandler struct {
	usersUC users.Usecase
}

func NewUsersHandler(usersUC users.Usecase) users.Handler {
	return &usersHandler{usersUC: usersUC}
}

func (h *usersHandler) Register(c *fiber.Ctx) error {
	registerUser := new(models.RegisterUser)

	if err := c.BodyParser(registerUser); err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":      "Bad Request",
			"status_code": fiber.StatusBadRequest,
			"message":     err.Error(),
			"result":      nil,
		})
	}

	res, err := h.usersUC.Register(registerUser)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":      "Internal Server Error",
			"status_code": fiber.StatusInternalServerError,
			"message":     err.Error(),
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

func (h *usersHandler) RemoveUser(c *fiber.Ctx) error {
	req := new(models.DeleteUser)
	userId := c.Locals("id").(string)
	if err := c.BodyParser(req); err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":      "Bad Request",
			"status_code": fiber.StatusBadRequest,
			"message":     nil,
			"result":      nil,
		})
	}
	req.Id = userId

	if err := h.usersUC.RemoveUser(req); err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":      "Internal Server Error",
			"status_code": fiber.StatusInternalServerError,
			"message":     "error, can't remove user.",
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
