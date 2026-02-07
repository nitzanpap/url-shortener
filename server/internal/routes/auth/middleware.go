package auth

import (
	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	service Service
}

func NewAuthMiddleware(service Service) *AuthMiddleware {
	return &AuthMiddleware{
		service: service,
	}
}

// RequireAuth is a middleware handler function that validates JWT tokens.
// It checks the httpOnly cookie first, then falls back to the Authorization header.
func (m *AuthMiddleware) RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := extractToken(c)
		if tokenString == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Authentication required"})
			return
		}

		userID, err := m.service.ValidateAndGetUserID(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid or expired token"})
			return
		}

		c.Set("user_id", userID)
		c.Next()
	}
}

func extractToken(c *gin.Context) string {
	// Check httpOnly cookie first
	if cookie, err := c.Cookie(CookieName); err == nil && cookie != "" {
		return cookie
	}

	// Fall back to Authorization header
	return extractBearerToken(c.GetHeader("Authorization"))
}

func extractBearerToken(header string) string {
	if len(header) > 7 && header[:7] == "Bearer " {
		return header[7:]
	}
	return ""
}
