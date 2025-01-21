package urls

import "github.com/gin-gonic/gin"

type handler struct {
	service *Service
}

func NewHandler(service *Service) Handler {
	return &handler{
		service: service,
	}
}

func (h *handler) ShortUrlHandler(c *gin.Context) {
	// Implementation here
}

func (h *handler) UrlGroupHandler(c *gin.Context) {
	// Implementation here
}
