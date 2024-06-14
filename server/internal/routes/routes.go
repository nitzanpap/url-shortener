package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/nitzanpap/url-shortener/server/internal/urls"
	"github.com/nitzanpap/url-shortener/server/pkg/utils"
)

func InitializeRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		utils.OkHandler(c, nil)
	})

	api := r.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			utils.OkHandler(c, nil)
		})
		v1 := api.Group("/v1")
		{
			v1.GET("/", func(c *gin.Context) {
				// This route should instead display the API documentation.
				// For now, it will redirect to the health check route.
				c.Redirect(http.StatusMovedPermanently, "/api/v1/health")
			})

			v1.GET("/health", func(c *gin.Context) {
				version := 1
				utils.OkHandler(c, &version)
			})
			
			urls.UrlsGroupHandler(v1)

		}
	}
}
