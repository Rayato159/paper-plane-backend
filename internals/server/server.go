package server

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/paper-plane/configs"
)

type Server struct {
	fiber *fiber.App
	db    *sqlx.DB
	cfg   *configs.Config
	file  *os.File
}

func NewServer(db *sqlx.DB, cfg *configs.Config, file *os.File) *Server {
	fiberConfig := configs.NewFiberConfig(cfg)
	return &Server{
		fiber: fiber.New(fiberConfig),
		db:    db,
		cfg:   cfg,
		file:  file,
	}
}

func (s *Server) Start() error {
	if err := s.MapHandlers(s.fiber); err != nil {
		return err
	}
	if err := s.fiber.Listen(fmt.Sprintf(":%s", s.cfg.Fiber.Port)); err != nil {
		return err
	}
	return nil
}
