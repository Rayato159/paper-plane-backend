package http

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/paper-plane/internals/auth"
	"github.com/paper-plane/internals/models"
)

type authHandler struct {
	authUC auth.Usecase
}

func NewAuthHandler(authUC auth.Usecase) auth.Handler {
	return &authHandler{authUC: authUC}
}

func (h *authHandler) Login(c *fiber.Ctx) error {
	credentials := new(models.Credentials)
	if err := c.BodyParser(credentials); err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":      "Bad Request",
			"status_code": fiber.StatusBadRequest,
			"message":     err.Error(),
			"result":      nil,
		})
	}

	res, err := h.authUC.Login(credentials)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":      "Unauthorized",
			"status_code": fiber.StatusUnauthorized,
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

func (h *authHandler) RefreshToken(c *fiber.Ctx) error {
	req := new(models.RefreshToken)
	if err := c.BodyParser(req); err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":      "Bad Request",
			"status_code": fiber.StatusBadRequest,
			"message":     err.Error(),
			"result":      nil,
		})
	}

	res, err := h.authUC.RefreshToken(req.RefreshToken)
	if err != nil {
		log.Println(err.Error())
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":      "Unauthorized",
			"status_code": fiber.StatusUnauthorized,
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
