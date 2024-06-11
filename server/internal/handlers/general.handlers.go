package handlers

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func HomeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func AboutHandler(c *gin.Context) {
    c.String(http.StatusOK, "About Page")
}
