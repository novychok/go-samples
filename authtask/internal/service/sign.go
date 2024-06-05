package service

import (
	"context"

	"github.com/novychok/go-samples/authtask/internal/entity"
)

type SignService interface {
	SignUp(ctx context.Context, signUp *entity.SignUp) (string, error)
	SignIn(ctx context.Context, signIn *entity.SignIn) (string, error)
	ParseAccessToken(accessToken string, registeredClaims *entity.RegisteredClaims, tokenSecret string) error
}
