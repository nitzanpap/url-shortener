package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"

	"github.com/nitzanpap/url-shortener/server/internal/configs"
	"github.com/nitzanpap/url-shortener/server/internal/routes/auth"
	"github.com/nitzanpap/url-shortener/server/internal/routes/urls"
	"github.com/nitzanpap/url-shortener/server/pkg/utils"
)

type Router struct {
	db                      *pgx.Conn
	jwtSecret               string
	urlHandler              urls.Handler
	supabaseAuthService     auth.SupabaseAuthService
	supabaseAuthMiddleware  *auth.SupabaseAuthMiddleware
}

func NewRouter(db *pgx.Conn, jwtSecret string, supabaseURL string) *Router {
	supabaseAuthService := auth.NewSupabaseAuthService(supabaseURL, jwtSecret)

	return &Router{
		db:                     db,
		jwtSecret:              jwtSecret,
		urlHandler:             urls.NewHandler(urls.NewService(urls.NewRepository(db))),
		supabaseAuthService:    supabaseAuthService,
		supabaseAuthMiddleware: auth.NewSupabaseAuthMiddleware(supabaseAuthService),
	}
}

func InitializeRoutes(r *gin.Engine, db *pgx.Conn, config *configs.Config) {
	r.GET("/", func(c *gin.Context) {
		utils.OkHandler(c, nil)
	})

	getUrlFromObfuscatedShortenedUrl := r.Group("/s")
	{
		urls.ShortUrlHandler(getUrlFromObfuscatedShortenedUrl, db)
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
				utils.TestDbConnection(db)
				c.JSON(http.StatusOK, gin.H{"status": "ok", "db": "ok"})
			})

			// Protected routes - require Supabase authentication
			protected := v1.Group("/")
			protected.Use(auth.NewSupabaseAuthMiddleware(auth.NewSupabaseAuthService(config.Supabase.URL, config.Supabase.JWTSecret)).RequireAuth())
			{
				urls.UrlGroupHandler(protected, db)
			}

			// Public routes - no authentication required
			public := v1.Group("/")
			{
				public.GET("/urls/public", func(c *gin.Context) {
					c.JSON(200, gin.H{"message": "This is a public endpoint"})
				})
			}
		}
	}
}


