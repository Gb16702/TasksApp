package routes

import (
	"regexp"
	"todoapp/core/utils"
	"todoapp/core/utils/validation"
	"todoapp/database"
	"todoapp/database/models"

	"github.com/alexedwards/argon2id"
	"github.com/gofiber/fiber/v2"
)

type body struct {
	Email 			string `json:"email"`
	Password 		string `json:"password"`
	PasswordConfirm string `json:"passwordConfirm"`
}

type verifyEmailBody struct {
	Email 			string `json:"email"`
	Token 			string `json:"token"`
}

 type loginBody struct {
 	Email 			string `json:"email"`
 	Password 		string `json:"password"`
 }

 const (
	ErrorMethod = "Méthode non autorisée"
	ErrorInvalid = "Requête invalide"
	ErrorInvalidIdentifiers = "Identifiants invalides"
	ErrorRequiredEmail = "L'email est requis"
	ErrorRequiredPassword = "Le mot de passe est requis"
 )

func HandleAuthRoutes(app *fiber.App) {
	app.Post("/api/auth/register", func (c *fiber.Ctx) error {

		if(c.Method() != fiber.MethodPost) {
			return utils.GenerateResponse(c, validation.GetErrorMessage("ErrorMethod"), 405)
		}

		var request body;
		if err := c.BodyParser(&request); err != nil {
			return utils.GenerateResponse(c, validation.GetErrorMessage("ErrorInvalid"), 400)
		}

		requiredFields := []struct{
			name 				string
			err 				string
		}{
			{name: "email", err: ErrorRequiredEmail},
			{name: "password", err: ErrorRequiredPassword},
		}

		pattern := regexp.MustCompile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`)

		for _, field := range requiredFields {
			switch field.name {
			case "email":
				isEmpty, err := validation.IsFieldEmpty(c, request.Email, ErrorRequiredEmail)
				if isEmpty {
					return err
				}

				if !pattern.MatchString(request.Email) {
					return utils.GenerateResponse(c, "L'email n'est pas valide", 400)
				}

			case "password":
				isEmpty, err := validation.IsFieldEmpty(c, request.Password, ErrorRequiredPassword)
				if isEmpty {
					return err
				}

				if request.Password != request.PasswordConfirm {
					return utils.GenerateResponse(c, "Les mots de passe ne correspondent pas", 400)
				}
			}
		}

		existingEmail := database.DB.Where("email = ?", request.Email).First(&models.User{})
		if existingEmail.RowsAffected != 0 {
			return utils.GenerateResponse(c, "Cette adresse mail est déjà utilisée", 400)
		}

		hashedPassword, err := argon2id.CreateHash(request.Password, argon2id.DefaultParams)
		if err != nil {
			return utils.GenerateResponse(c, "Erreur lors du hashage du mot de passe", 500)
		}

		user := models.User{
			Email: request.Email,
			Password: string(hashedPassword),
		}

		result := database.DB.Create(&user)
		if result.Error != nil {
			return utils.GenerateResponse(c, "Erreur lors de la création du compte", 500)
		}

		token, err := utils.GenerateToken(request.Email);
		if err != nil {
			return utils.GenerateResponse(c, "Erreur lors de la génération du token", 500)
		} else {
			error := utils.SendVerificationEmail(user,request.Email, token)

			if error != nil {
				return utils.GenerateResponse(c, "Erreur lors de l'envoi de l'email de vérification", 500)
			}
		}

		return utils.GenerateResponse(c, "Compte créé avec succès", 200)
	})

	app.Post("/api/auth/login", func(c *fiber.Ctx) error {

		if(c.Method() != fiber.MethodPost) {
			return utils.GenerateResponse(c, validation.GetErrorMessage("ErrorMethod"), 405)
		}

		headers := c.Get("Authorization")
		if headers != "" {
			token, err := utils.ExtractToken(headers)
			if err != nil {
				return utils.GenerateResponse(c, "Erreur lors de l'extraction du token", 500)
			}

			decodedToken, err := utils.DecodeToken(token)
			if err != nil {
				return utils.GenerateResponse(c, "Erreur lors du décodage du token", 500)
			}

			var user models.User
			result := database.DB.Where("email = ?", decodedToken).Select("id").First(&user);
			if result.Error != nil {
				return utils.GenerateResponse(c, "Erreur serveur", 500)
			}

			session, err := utils.GenerateSessionToken(decodedToken, user.ID)
			if err != nil {
				return utils.GenerateResponse(c, "Erreur lors de la génération du token de session", 500);
			}
			return utils.GenerateResponse(c, "Connexion réussie", 200, session);

		} else {

			var request loginBody;
			if err := c.BodyParser(&request); err != nil {
				return utils.GenerateResponse(c, validation.GetErrorMessage("ErrorInvalid"), 400)
			}

			requiredFields := []string{"email", "password"}
			for _, field := range requiredFields {
				switch field {
				case "email":
					isEmpty, err := validation.IsFieldEmpty(c, request.Email, ErrorRequiredEmail)
					if isEmpty {
						return err
					}

				case "password":
					isEmpty, err := validation.IsFieldEmpty(c, request.Password, ErrorRequiredPassword)
					if isEmpty {
						return err
					}
				}
			}

			var user models.User
			result := database.DB.Where("email = ?", request.Email).First(&user);
			if result.Error != nil {
				return utils.GenerateResponse(c, validation.GetErrorMessage("ErrorInvalidIdentifiers"), 500)
			}

			match, _, err := argon2id.CheckHash(request.Password, user.Password)
			if err != nil || !match {
				return utils.GenerateResponse(c, validation.GetErrorMessage("ErrorInvalidIdentifiers"), 500)
			}

			// if(!user.Verified) {
			// 	return utils.GenerateResponse(c, "Adresse mail non vérifiée", 400)
			// }

			session, err := utils.GenerateSessionToken(request.Email, user.ID)
			if err != nil {
				return utils.GenerateResponse(c, "Erreur lors de la génération du token de session", 500)
			}

			return c.Status(200).JSON(fiber.Map{
				"message": "Connexion réussie",
				"session": session,
				"ID" : user.ID,
			})
		}
	})

	app.Post("/api/auth/verify-email", func (c *fiber.Ctx) error {
		if(c.Method() != fiber.MethodPost) {
			return utils.GenerateResponse(c, validation.GetErrorMessage("ErrorMethod"), 405)
		}

		var request verifyEmailBody;
		if err := c.BodyParser(&request); err != nil {
			return utils.GenerateResponse(c, validation.GetErrorMessage("ErrorInvalid"), 400)
		}

		requiredFields := []string{"user", "token"}
		for _, field := range requiredFields {
			switch field {
			case "user":
				isEmpty, err := validation.IsFieldEmpty(c, request.Email, ErrorRequiredEmail)
				if isEmpty {
					return err
				}
		 	}
		}

		verifiedUser := database.DB.Where("email = ?", request.Email).First(&models.User{}).Update("verified", true)
		if verifiedUser.Error != nil {
			return utils.GenerateResponse(c, "Erreur lors de la vérification de l'email", 500)
		}

		return utils.GenerateResponse(c, "Email vérifié avec succès", 200)
	})
}
