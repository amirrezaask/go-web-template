package auth

import (
	"app/config"
	"errors"

	"github.com/dgrijalva/jwt-go"
)

var ErrInvalidJWT = errors.New("invalid jwt token")

func NewJWTToken(claims jwt.Claims) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := config.Config.Auth.JWT.Secret
	return t.SignedString(secret)
}

func ValidateToken(token string) (*jwt.Token, error) {
	secret := config.Config.Auth.JWT.Secret
	t, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return nil, err
	}
	if !t.Valid {
		return nil, ErrInvalidJWT
	}
	return t, nil
}

type Identity interface {
	//unique identifier for this user, probably database primary key
	ID() int
	//whether this user has access to given key, it can be a path, resource or anything
	HasAccessTo(key string) error
}

//IdentityProvider is function type that gets a jwt token and creates the identity based on that
type IdentityProvider func(*jwt.Token) (Identity, error)
