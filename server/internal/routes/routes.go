package routes

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"

	"github.com/nitzanpap/url-shortener/server/internal/routes/auth"
	"github.com/nitzanpap/url-shortener/server/internal/routes/urls"
	"github.com/nitzanpap/url-shortener/server/internal/routes/users"
	"github.com/nitzanpap/url-shortener/server/pkg/utils"
)

type Router struct {
	db             *pgx.Conn
	jwtSecret      string
	urlHandler     urls.Handler
	authHandler    auth.Handler
	authMiddleware *auth.AuthMiddleware
}

func NewRouter(db *pgx.Conn, jwtSecret string) *Router {
	userRepo := users.NewUserRepository(db)
	authService := auth.NewService(userRepo, jwtSecret, 24*time.Hour)

	return &Router{
		db:             db,
		jwtSecret:      jwtSecret,
		urlHandler:     urls.NewHandler(urls.NewService(urls.NewRepository(db))),
		authHandler:    auth.NewHandler(authService),
		authMiddleware: auth.NewAuthMiddleware(authService),
	}
}

func InitializeRoutes(r *gin.Engine, db *pgx.Conn) {
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
				// This route should instead display the API documentation.
				// For now, it will redirect to the health check route.
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

			urls.URLGroupHandler(v1, db)
		}
	}

	authRouter := NewRouter(db, "your_jwt_secret")
	authHandler := authRouter.Setup()
	r.Any("/auth/*path", func(c *gin.Context) {
		authHandler.ServeHTTP(c.Writer, c.Request)
	})
}

func (r *Router) Setup() http.Handler {
	router := gin.New()

	// Add your auth routes here
	auth := router.Group("/")
	{
		auth.POST("/login", r.authHandler.Login)
		auth.POST("/register", r.authHandler.Register)
		// Add other auth routes as needed
	}

	return router
}
