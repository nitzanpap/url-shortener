package auth

import (
	"github.com/gin-gonic/gin"
)

type credentials struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type Handler interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
}

type handler struct {
	service Service
}

func NewHandler(service Service) Handler {
	return &handler{service: service}
}

func (h *handler) Register(c *gin.Context) {
	var creds Credentials
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	if err := h.service.Register(creds.Email, creds.Password); err != nil {
		statusCode := 500
		if err == ErrInvalidCredentials {
			statusCode = 400
		}
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	c.Status(201)
}

func (h *handler) Login(c *gin.Context) {
	var creds Credentials
	if err := c.ShouldBindJSON(&creds); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	token, err := h.service.Login(creds.Email, creds.Password)
	if err != nil {
		statusCode := 401
		if err == ErrUserNotFound {
			statusCode = 404
		}
		c.JSON(statusCode, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, LoginResponse{Token: token})
}
