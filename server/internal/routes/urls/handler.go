package urls

import "github.com/gin-gonic/gin"

// Handler defines the interface for URL handling operations
type Handler interface {
	// Add methods that your URL handler needs to implement
	// For example:
	ShortURLHandler(c *gin.Context)
	URLGroupHandler(c *gin.Context)
}
