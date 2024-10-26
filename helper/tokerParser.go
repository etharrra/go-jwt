package helper

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func TokenParse(tokenString string) (*jwt.Token, error) {
	secretKey := getSecretKey()
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return secretKey, nil
	})
}

func getSecretKey() []byte {
	return []byte(os.Getenv("SECRECT"))
}
