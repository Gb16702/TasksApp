package utils

import (
	"strings"
)

func ExtractToken(bearer string) (string, error) {

	if len(bearer) == 0 {
		return "", nil
	}

	splitToken := strings.Split(bearer, " ")

	if len(splitToken) != 2 {
		return "", nil
	}

	return splitToken[1], nil
}
