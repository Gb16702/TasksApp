package utils

import (
	"fmt"
	"os"
	"time"
	"github.com/golang-jwt/jwt/v5"
)

func DecodeToken(tokenString string) (string, error) {

	key := []byte(os.Getenv("JWT_KEY"));

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if claims["authorized"] != true {
		return "", fmt.Errorf("vous n'êtes pas autorisé à accéder à cette ressource")
	}

	if !ok {
		return "", fmt.Errorf("le token n'est pas valide")
	}

	expirationDate := time.Unix(int64(claims["exp"].(float64)), 0)

	if time.Now().After(expirationDate) {
		return "", fmt.Errorf("le token a expiré")
	}

	user := claims["user"].(string)

	return user, nil
}
