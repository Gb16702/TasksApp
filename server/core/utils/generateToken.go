package utils

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)


func GenerateToken(email string) (string, error) {

	var key = []byte(os.Getenv("JWT_KEY"));

	token := jwt.New(jwt.SigningMethodHS256);

	claims := token.Claims.(jwt.MapClaims);

	claims["exp"] = time.Now().Add(time.Hour).Unix();
	claims["authorized"] = true;
	claims["user"] = email;

	tokenString, err := token.SignedString(key);

	if err != nil {
		return "", err;
	}

	return tokenString, nil;
}

func GenerateSessionToken(email string, ID uint) (string, error) {
	var key = []byte(os.Getenv("JWT_KEY"));

	token := jwt.New(jwt.SigningMethodHS256);

	claims := token.Claims.(jwt.MapClaims);

	claims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix();
	claims["authorized"] = true;
	claims["user"] = email;
	claims["id"] = ID;

	tokenString, err := token.SignedString(key);

	if err != nil {
		return "", err;
	}

	return tokenString, nil;
}
