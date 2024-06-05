package sign

import (
	"context"
	"errors"
	"log/slog"
	"os"
	"strconv"
	"time"

	"github.com/novychok/go-samples/authtask/internal/entity"
	"github.com/novychok/go-samples/authtask/internal/repository"
	"github.com/novychok/go-samples/authtask/internal/service"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type srv struct {
	repo repository.SignRepository

	l *slog.Logger
}

func (s *srv) SignUp(ctx context.Context, signUp *entity.SignUp) (string, error) {
	l := s.l.With(slog.String("method", "signUp"))

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(signUp.Password), 10)
	if err != nil {
		return "", err
	}

	user := &entity.User{
		Username:     signUp.Username,
		PasswordHash: string(hashedPassword),
	}

	err = s.repo.Create(ctx, user)
	if err != nil {
		l.ErrorContext(ctx, "Failed to save user", "err", err)
		return "", err
	}

	jwtToken, err := s.jwtWithClaims(user, l)
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}

func (s *srv) SignIn(ctx context.Context, signIn *entity.SignIn) (string, error) {
	l := s.l.With(slog.String("method", "signIn"))

	user, err := s.repo.Get(ctx, signIn.Username)
	if err != nil {
		l.ErrorContext(ctx, "Failed to find user by email", "err", err)
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(signIn.Password), []byte(user.PasswordHash)); err != nil {
		l.ErrorContext(ctx, "Failed while compare hash password", "err", errors.New("invalid password"))
		return "", errors.New("invalid password")
	}

	jwtToken, err := s.jwtWithClaims(user, l)
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}

func (s *srv) jwtWithClaims(user *entity.User, l *slog.Logger) (string, error) {
	timeNow := time.Now()
	tokenExpiration, err := time.ParseDuration(os.Getenv("TOKEN_EXPIRATION") + os.Getenv("TOKEN_MIN"))
	if err != nil {
		return "", err
	}

	claims := entity.RegisteredClaims{
		Name: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: timeNow.Add(tokenExpiration).Unix(),
			Id:        strconv.Itoa(user.ID),
			IssuedAt:  timeNow.Unix(),
			NotBefore: timeNow.Unix(),
		},
	}

	jwtToken, err := newAccessToken(&claims, os.Getenv("TOKEN_SECRET"))
	if err != nil {
		l.Error("Failed to create access token", "err", err)
		return "", err
	}

	return jwtToken, nil
}

func newAccessToken(registeredClaims *entity.RegisteredClaims, tokenSecret string) (string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, registeredClaims)
	return accessToken.SignedString([]byte(tokenSecret))
}

func (s *srv) ParseAccessToken(accessToken string, registeredClaims *entity.RegisteredClaims, tokenSecret string) error {
	parsedToken, err := jwt.ParseWithClaims(accessToken, registeredClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenSecret), nil
	})
	if err != nil {
		return err
	}

	if !parsedToken.Valid {
		return entity.ErrInvalidToken
	}

	return nil
}

func New(repo repository.SignRepository, l *slog.Logger) service.SignService {
	return &srv{
		repo: repo,
		l:    l,
	}
}
