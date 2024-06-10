package main

import (
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func uploadFile(c echo.Context) error {
	// Read form file
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	// Source
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(file.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.String(http.StatusOK, "File uploaded successfully")
}

func main() {
	e := echo.New()

	// Routes
	e.POST("/upload", uploadFile)
	e.GET("/upload-status", func(c echo.Context) error {
		// Dummy endpoint for upload status
		return c.String(http.StatusOK, "Upload status: Success")
	})

	// Start server
	e.Start(":8080")
}
