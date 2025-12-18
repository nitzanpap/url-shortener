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

func (h *handler) ShortURLHandler(c *gin.Context) {
	// Implementation here
}

func (h *handler) URLGroupHandler(c *gin.Context) {
	// Implementation here
}
