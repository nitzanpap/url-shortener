package serverconfigs

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nitzanpap/url-shortener/server/internal/configs"
)

// GetGinMode returns the gin mode based on the environment
func GetGinMode(config *configs.Config) string {
	if config.Environment == configs.Local {
		return gin.DebugMode
	}
	return gin.ReleaseMode
}

// SetupGinServer sets up the gin server with the appropriate configuration
func SetupGinServer(config *configs.Config) *gin.Engine {
	// Use gin.New() instead of gin.Default() to avoid duplicate middleware warning
	router := gin.New()

	// Manually add the middleware we need
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// In local we don't need to set trusted proxies. In all other environments, we should set the trusted proxies
	// ! TODO: Set trusted proxies appropriately
	router.SetTrustedProxies([]string{})

	// Add CORS middleware with sanitized client origin
	router.Use(setupCORSMiddleware(sanitizeOrigin(config.ClientOrigin)))

	return router
}

// sanitizeOrigin ensures the origin doesn't have a trailing slash
func sanitizeOrigin(origin string) string {
	return strings.TrimSuffix(origin, "/")
}

// setupCORSMiddleware creates middleware for handling CORS
func setupCORSMiddleware(clientOrigin string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the Origin header from the request
		requestOrigin := c.Request.Header.Get("Origin")

		// If the request has an Origin header and it matches our allowed origin
		// or if we want to allow all origins with "*"
		if requestOrigin != "" && (clientOrigin == "*" || requestOrigin == clientOrigin) {
			c.Writer.Header().Set("Access-Control-Allow-Origin", requestOrigin)
		} else {
			c.Writer.Header().Set("Access-Control-Allow-Origin", clientOrigin)
		}

		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
