package routes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"

	"github.com/nitzanpap/url-shortener/server/internal/configs"
	"github.com/nitzanpap/url-shortener/server/internal/middleware"
	"github.com/nitzanpap/url-shortener/server/internal/routes/auth"
	"github.com/nitzanpap/url-shortener/server/internal/routes/urls"
	"github.com/nitzanpap/url-shortener/server/internal/routes/users"
	"github.com/nitzanpap/url-shortener/server/pkg/utils"
)

type Router struct {
	db             *pgx.Conn
	jwtSecret      string
	environment    configs.Environment
	urlHandler     urls.Handler
	authHandler    auth.Handler
	authMiddleware *auth.AuthMiddleware
}

func NewRouter(db *pgx.Conn, jwtSecret string, environment configs.Environment) *Router {
	userRepo := users.NewUserRepository(db)
	authService := auth.NewService(userRepo, jwtSecret, 24*time.Hour)

	return &Router{
		db:             db,
		jwtSecret:      jwtSecret,
		environment:    environment,
		urlHandler:     urls.NewHandler(urls.NewService(urls.NewRepository(db))),
		authHandler:    auth.NewHandler(authService, environment),
		authMiddleware: auth.NewAuthMiddleware(authService),
	}
}

func InitializeRoutes(r *gin.Engine, db *pgx.Conn, config *configs.Config) {
	authRateLimiter := middleware.NewRateLimiter(10, 5)
	urlRateLimiter := middleware.NewRateLimiter(30, 10)

	r.GET("/", func(c *gin.Context) {
		utils.OkHandler(c, nil)
	})

	getURLFromObfuscatedShortenedURL := r.Group("/s")
	{
		urls.ShortURLHandler(getURLFromObfuscatedShortenedURL, db)
	}

	api := r.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			utils.OkHandler(c, nil)
		})
		v1 := api.Group("/v1")
		{
			v1.GET("/", func(c *gin.Context) {
				c.Redirect(http.StatusMovedPermanently, "/api/v1/health")
			})

			v1.GET("/health", func(c *gin.Context) {
				version := 1
				utils.OkHandler(c, &version)
			})

			v1.GET("/db/health", func(c *gin.Context) {
				utils.TestDBConnection(db)
				c.JSON(http.StatusOK, gin.H{"status": "ok", "db": "ok"})
			})

			rateLimitedV1 := v1.Group("/")
			rateLimitedV1.Use(urlRateLimiter.Middleware())
			urls.URLGroupHandler(rateLimitedV1, db)
		}
	}

	authRouter := NewRouter(db, config.JWTSecret, config.Environment)
	authHandler := authRouter.Setup(authRateLimiter)
	r.Any("/auth/*path", func(c *gin.Context) {
		authHandler.ServeHTTP(c.Writer, c.Request)
	})
}

func (r *Router) Setup(rateLimiter *middleware.RateLimiter) http.Handler {
	router := gin.New()

	authGroup := router.Group("/")
	authGroup.Use(rateLimiter.Middleware())
	{
		authGroup.POST("/login", r.authHandler.Login)
		authGroup.POST("/register", r.authHandler.Register)
		authGroup.POST("/logout", r.authHandler.Logout)
	}

	return router
}
