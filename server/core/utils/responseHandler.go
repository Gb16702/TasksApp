package utils

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

func GenerateResponse(c *fiber.Ctx, message string, statusCode int, optionalData ...interface{}) error {
	response := fiber.Map{
		"message": message,
	}

	for i := 0; i < len(optionalData); i += 2 {
		if i+1 < len(optionalData) {
			key, ok := optionalData[i].(string)
			if !ok {
				return errors.New("invalid optional data format, key must be a string")
			}
			response[key] = optionalData[i+1]
		}
	}

	return c.Status(statusCode).JSON(response)
}
