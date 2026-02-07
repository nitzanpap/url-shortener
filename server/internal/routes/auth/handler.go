package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nitzanpap/url-shortener/server/internal/configs"
)

type Handler interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	Logout(c *gin.Context)
}

type handler struct {
	service     Service
	environment configs.Environment
}

func NewHandler(service Service, environment configs.Environment) Handler {
	return &handler{service: service, environment: environment}
}

func (h *handler) Register(c *gin.Context) {
	var creds Credentials
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := h.service.Register(creds.Email, creds.Password); err != nil {
		if err == ErrInvalidCredentials {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid credentials"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Registration failed"})
		return
	}

	c.Status(http.StatusCreated)
}

func (h *handler) Login(c *gin.Context) {
	var creds Credentials
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	token, err := h.service.Login(creds.Email, creds.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	secure := h.environment == configs.Production
	sameSite := http.SameSiteLaxMode
	if secure {
		sameSite = http.SameSiteStrictMode
	}
	c.SetSameSite(sameSite)
	c.SetCookie(CookieName, token, CookieMaxAge, CookiePath, "", secure, CookieHTTPOnly)

	c.JSON(http.StatusOK, LoginResponse{Token: token})
}

func (h *handler) Logout(c *gin.Context) {
	secure := h.environment == configs.Production
	sameSite := http.SameSiteLaxMode
	if secure {
		sameSite = http.SameSiteStrictMode
	}
	c.SetSameSite(sameSite)
	c.SetCookie(CookieName, "", -1, CookiePath, "", secure, CookieHTTPOnly)

	c.JSON(http.StatusOK, gin.H{"message": "Logged out"})
}
