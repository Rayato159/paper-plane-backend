package server

import (
	"log"

	usersHttp "github.com/paper-plane/internals/users/deliveries/http"
	usersRepository "github.com/paper-plane/internals/users/repositories"
	usersUsecase "github.com/paper-plane/internals/users/usecases"

	authHttp "github.com/paper-plane/internals/auth/deliveries/http"
	authRepository "github.com/paper-plane/internals/auth/repositories"
	authUsecase "github.com/paper-plane/internals/auth/usecases"

	accountsHttp "github.com/paper-plane/internals/accounts/deliveries/http"
	accountsRepository "github.com/paper-plane/internals/accounts/repositories"
	accountsUsecase "github.com/paper-plane/internals/accounts/usecases"

	transactionsHttp "github.com/paper-plane/internals/transactions/deliveries/http"
	transactionsRepository "github.com/paper-plane/internals/transactions/repositories"
	transactionsUsecase "github.com/paper-plane/internals/transactions/usecases"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func (s *Server) MapHandlers(a *fiber.App) error {
	// For reccords
	a.Use(logger.New(logger.Config{
		Format:     "${time} | ${pid} | ${ip} | ${status} | ${method} ${path}\n",
		TimeFormat: "2006-01-02T15:04:05",
		TimeZone:   "Thailand/Bangkok",
		Output:     s.file,
	}))

	// For console
	a.Use(logger.New(logger.Config{
		Format:     "${time} | ${pid} | ${ip} | ${status} | ${method} ${path}\n",
		TimeFormat: "2006-01-02T15:04:05",
		TimeZone:   "Thailand/Bangkok",
	}))

	v1 := a.Group("/v1")

	authGroup := v1.Group("/auth")
	usersRepo := usersRepository.NewUsersRepository(s.db)
	authRepo := authRepository.NewAuthRepository(s.db)
	authUsecase := authUsecase.NewAuthUsecase(authRepo)
	authHandler := authHttp.NewAuthHandler(authUsecase)
	authHttp.MapAuthRoute(authGroup, authHandler)

	accountsGroup := v1.Group("/accounts")
	accountsRepo := accountsRepository.NewAccountsRepository(s.db)
	usersUsecase := usersUsecase.NewUsersUsecase(usersRepo, accountsRepo)
	accountsUsecase := accountsUsecase.NewAccountsUsecase(accountsRepo)
	accountsHandler := accountsHttp.NewAccountsHandler(accountsUsecase)
	accountsHttp.MapAccountsRoute(accountsGroup, accountsHandler, usersUsecase)

	usersGroup := v1.Group("/users")
	usersHandler := usersHttp.NewUsersHandler(usersUsecase)
	usersHttp.MapUsersRoute(usersGroup, usersHandler, usersUsecase)

	transactionsGroup := v1.Group("/transactions")
	transactionsRepo := transactionsRepository.NewTransactionsRepository(s.db)
	transactionsUsecase := transactionsUsecase.NewTransactionsUsecase(transactionsRepo, accountsRepo, usersRepo)
	transactionsHandler := transactionsHttp.NewTransactionsHandler(transactionsUsecase)
	transactionsHttp.MapTransactionsRoute(transactionsGroup, transactionsHandler, usersUsecase)

	a.Use(func(c *fiber.Ctx) error {
		log.Println("sorry, endpoint is not found")
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":      "Not Found",
			"status_code": fiber.StatusNotFound,
			"message":     "sorry, endpoint is not found",
		})
	})
	return nil
}
