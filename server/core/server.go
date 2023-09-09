package server

import (
	"github.com/gofiber/fiber/v2"
)

func HandleServerStart(port, DB_URL string) {

	app := fiber.New();

	app.Get("/",  func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!");
	})

	app.Listen(port)
}
