package utils

import (
	"github.com/gofiber/fiber/v2"
)

func GenerateResponse(c *fiber.Ctx, message string, statusCode int, optionalData ...interface{}) error {
	response := fiber.Map{
		"message": message,
	}

	if len(optionalData) > 0 {
		response["data"] = optionalData[0]
	}

	return c.Status(statusCode).JSON(response)
}
