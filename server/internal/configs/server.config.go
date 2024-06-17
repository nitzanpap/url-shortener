package configs

import (
	"log"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/nitzanpap/url-shortener/server/pkg/colors"
)

func SetupGinServer(config *Config) *gin.Engine {
	r := gin.Default()
	SetCors(r, config)
	return r
}
func SetCors(r *gin.Engine, config *Config) {
	r.Use(cors.New(cors.Config{
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			isValidOrigin := strings.Contains(config.ClientOrigin, origin) || strings.HasPrefix(origin, "http://localhost:")
			if !isValidOrigin {
				log.Printf(colors.Error("Invalid origin: %s"), origin)
			}
			return isValidOrigin
		},
		MaxAge: 12 * time.Hour,
	}))
}

func GetGinMode(config *Config) (mode string) {
	switch config.Environment {
	case Development:
		return gin.DebugMode
	case Production:
		return gin.ReleaseMode
	default:
		log.Fatalf(colors.Error("Invalid environment: %s"), config.Environment)
		return gin.DebugMode
	}
}
