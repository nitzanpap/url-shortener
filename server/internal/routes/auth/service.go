package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/nitzanpap/url-shortener/server/internal/models"
	"github.com/nitzanpap/url-shortener/server/internal/routes/users"
)

type Service interface {
	Register(email, password string) error
	Login(email, password string) (string, error)
	ValidateAndGetUserID(tokenString string) (uint, error)
}

type service struct {
	userRepo    users.UserRepository
	jwtSecret   string
	tokenExpiry time.Duration
}

func NewService(userRepo users.UserRepository, jwtSecret string, tokenExpiry time.Duration) Service {
	return &service{
		userRepo:    userRepo,
		jwtSecret:   jwtSecret,
		tokenExpiry: tokenExpiry,
	}
}

func (s *service) Register(email, password string) error {
	user := &models.User{
		Email:    email,
		Password: password,
	}

	if err := user.HashPassword(); err != nil {
		return err
	}

	return s.userRepo.Create(user)
}

func (s *service) Login(email, password string) (string, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return "", ErrUserNotFound
	}

	if err := user.ComparePassword(password); err != nil {
		return "", ErrInvalidCredentials
	}

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(s.tokenExpiry).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *service) ValidateAndGetUserID(tokenString string) (uint, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(s.jwtSecret), nil
	})

	if err != nil || !token.Valid {
		return 0, ErrInvalidToken
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, ErrInvalidClaims
	}

	userID, ok := claims["user_id"].(float64)
	if !ok {
		return 0, ErrInvalidClaims
	}

	return uint(userID), nil
}
