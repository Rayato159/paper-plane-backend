package configs

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func NewFiberConfig(c *Config) fiber.Config {
	readTimeoutSecondCount, _ := strconv.Atoi(c.Fiber.ServerReadTimeOut)

	return fiber.Config{
		ReadTimeout: time.Second * time.Duration(readTimeoutSecondCount),
		BodyLimit:   10 * 1024 * 1024,
	}
}
