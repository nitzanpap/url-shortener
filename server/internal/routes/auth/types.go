package auth

import (
	"errors"
	"time"
)

// Credentials represents the user authentication request payload.
type Credentials struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,max=72"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

// Cookie configuration
const (
	CookieName     = "auth_token"
	CookieMaxAge   = int(24 * time.Hour / time.Second) // 86400 seconds
	CookiePath     = "/"
	CookieHTTPOnly = true
)

// Error definitions
var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidToken       = errors.New("invalid token")
	ErrInvalidClaims      = errors.New("invalid token claims")
)
