package routes

import (
	"github.com/gofiber/fiber/v2"
)

type body struct {
	Email 		string `json:"email"`
	Password 	string `json:"password"`
}

func HandleAuthRoutes(app *fiber.App) {
	app.Post("/api/auth/register", func (c *fiber.Ctx) error {

		if(c.Method() != fiber.MethodPost) {
			return c.Status(405).JSON(fiber.Map{
				"message": "Méthode non autorisée",
			})
		}

		var request body;

		err := c.BodyParser(&request);

		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": "Requête",
			});
		}

		// ...

		return c.Status(200).JSON(fiber.Map{
			"message" : "Compte créé avec succès",
		})
	})
}
