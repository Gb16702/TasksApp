package validation

const (
	ErrorMethod              = "Méthode non autorisée"
	ErrorInvalid             = "Requête invalide"
	ErrorInvalidIdentifiers  = "Identifiants invalides"
	ErrorRequiredEmail     	 = "L'email est requis"
	ErrorRequiredPassword    = "Le mot de passe est requis"
	ErrorInvalidParams       = "Paramètres invalides"
)

func GetErrorMessage(code string) string {
	switch code {
	case "ErrorMethod":
		return ErrorMethod
	case "ErrorInvalid":
		return ErrorInvalid
	case "ErrorInvalidIdentifiers":
		return ErrorInvalidIdentifiers
	case "ErrorRequiredEmail":
		return ErrorRequiredEmail
	case "ErrorRequiredPassword":
		return ErrorRequiredPassword
	case "ErrorInvalidParams":
		return "Paramètres invalides"
	default:
		return "Erreur inconnue"
	}
}
