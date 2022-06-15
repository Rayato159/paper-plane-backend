package middlewares

import (
	"errors"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/paper-plane/internals/models"
	"github.com/paper-plane/internals/users"
)

func JwtAuthentication(u users.Usecase) fiber.Handler {
	return func(c *fiber.Ctx) error {
		accessToken := strings.TrimPrefix(c.Get("Authorization"), "Bearer ")
		if accessToken == "" {
			log.Println(errors.New("error, authorization header is empty."))
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status":      "Unauthorized",
				"status_code": fiber.StatusUnauthorized,
				"message":     "error, unauthorized",
				"result":      nil,
			})
		}

		token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
		})
		if err != nil {
			log.Println(errors.New("error, unexpected signing method."))
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status":      "Unauthorized",
				"status_code": fiber.StatusUnauthorized,
				"message":     "error, unauthorized",
				"result":      nil,
			})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			// Payload Check!!!

			payload := new(models.UserJWTPayload)
			payload.UserId = claims["id"].(string)
			payload.Username = claims["username"].(string)
			payload.Expires = claims["exp"].(float64)

			tokenExpiresChan := make(chan bool)
			userPayloadJWTCheckChan := make(chan error)

			go func() {
				tokenExpires := IsTokenExpires(payload.Expires)
				tokenExpiresChan <- tokenExpires
				close(tokenExpiresChan)
			}()

			go func() {
				err := u.UserJWTPayloadCheck(payload)
				userPayloadJWTCheckChan <- err
				close(userPayloadJWTCheckChan)
			}()

			tokenExpires := <-tokenExpiresChan
			userPayloadJWTCheck := <-userPayloadJWTCheckChan

			if tokenExpires {
				log.Println(errors.New("error, token was expired."))
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"status":      "Unauthorized",
					"status_code": fiber.StatusUnauthorized,
					"message":     "error, unauthorized",
					"result":      nil,
				})
			}
			if userPayloadJWTCheck != nil {
				log.Println(errors.New("error, users not found."))
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"status":      "Unauthorized",
					"status_code": fiber.StatusUnauthorized,
					"message":     "error, unauthorized",
					"result":      nil,
				})
			}

			c.Locals("id", payload.UserId)
			c.Locals("username", payload.Username)

			return c.Next()
		}

		log.Println(errors.New("error, signature is invalid."))
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":      "Unauthorized",
			"status_code": fiber.StatusUnauthorized,
			"message":     "error, unauthorized",
			"result":      nil,
		})
	}
}

func IsTokenExpires(t float64) bool {
	if int64(t) > time.Now().Unix()-(86400*14) {
		return false
	}
	return true
}
