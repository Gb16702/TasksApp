package routes

import (
	"fmt"

	"regexp"
	"todoapp/database"
	"todoapp/database/models"
	"todoapp/core/utils"

	"github.com/alexedwards/argon2id"
	"github.com/gofiber/fiber/v2"

)

type body struct {
	Email 			string `json:"email"`
	Password 		string `json:"password"`
	PasswordConfirm string `json:"passwordConfirm"`
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

		fmt.Println(request, err)

		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": "Requête",
			});
		}

		requiredFields := []string{"email", "password", "passwordConfirm"}

		pattern := regexp.MustCompile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`)

		for _, field := range requiredFields {
			switch field {
			case "email":
				if request.Email == "" {
					return c.Status(400).JSON(fiber.Map{
						"message": "L'email est requis",
					})
				} else if !pattern.MatchString(request.Email) {
					return c.Status(400).JSON(fiber.Map{
						"message": "L'email n'est pas valide",
					})
				}

			case "password":
				if request.Password == "" {
					return c.Status(400).JSON(fiber.Map{
						"message": "Le mot de passe est requis",
					})
				} else if request.Password != request.PasswordConfirm {
					return c.Status(400).JSON(fiber.Map{
						"message": "Les mots de passe ne correspondent pas",
					})
				}
			}
		}

		existingEmail := database.DB.Where("email = ?", request.Email).First(&models.User{})

		if existingEmail.RowsAffected != 0 {
			return c.Status(400).JSON(fiber.Map{
				"message": "L'email existe déjà",
			})
		}

		hashedPassword, err := argon2id.CreateHash(request.Password, argon2id.DefaultParams)

		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": "Erreur lors de la création du mot de passe",
			})
		}

		user := models.User{
			Email: request.Email,
			Password: string(hashedPassword),
		}

		result := database.DB.Create(&user)

		if result.Error != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": "Erreur lors de la création du compte",
			})
		}


		error := utils.SendVerificationEmail(user,request.Email)

		if error != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": "Erreur lors de l'envoi de l'email de vérification",
			})
		}

		return c.Status(200).JSON(fiber.Map{
			"message" : "Compte créé avec succès",
		})
	})
}
