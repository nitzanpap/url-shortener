package auth

import (
	"github.com/gin-gonic/gin"
)

type SupabaseAuthMiddleware struct {
	service SupabaseAuthService
}

func NewSupabaseAuthMiddleware(service SupabaseAuthService) *SupabaseAuthMiddleware {
	return &SupabaseAuthMiddleware{
		service: service,
	}
}

// RequireAuth middleware for protecting routes with Supabase authentication
func (m *SupabaseAuthMiddleware) RequireAuth() gin.HandlerFunc {
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

		// Validate the token with Supabase
		user, err := m.service.ValidateToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid or expired token"})
			return
		}

		// Set user information in the context
		c.Set("user_id", user.ID)
		c.Set("user_email", user.Email)
		c.Set("user_role", user.Role)
		c.Set("user", user)
		c.Next()
	}
}

