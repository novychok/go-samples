package entity

import (
	"errors"

	"github.com/golang-jwt/jwt"
)

var (
	ErrSetCookieResponse = errors.New("error set cookie to response")
	ErrInvalidToken      = errors.New("invalid token")
)

type SignUp struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SignIn struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisteredClaims struct {
	Name string
	jwt.StandardClaims
}
