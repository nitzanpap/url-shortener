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
func (m *AuthMiddleware) RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Authorization header required"})
			return
		}

		// Extract Bearer token
		tokenString := extractBearerToken(authHeader)
		if tokenString == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid token format"})
			return
		}

		userID, err := m.service.ValidateAndGetUserID(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
			return
		}

		c.Set("user_id", userID)
		c.Next()
	}
}

func extractBearerToken(header string) string {
	if len(header) > 7 && header[:7] == "Bearer " {
		return header[7:]
	}
	return ""
}
