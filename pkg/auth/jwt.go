package auth

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

func VerifyJWTToken(key string, t string) error {
	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	if err == nil && token.Valid {
		return nil
	}
	return fmt.Errorf("jwt is not valid: %w", err)
}
func NewJWTToken(signKey string, claims jwt.Claims) (string, error) {
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	token.Claims = claims
	return token.SignedString(signKey)
}
