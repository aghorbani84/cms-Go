package handlers

import (
	"ginwebapp/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Hello handles the root endpoint
func Hello(c echo.Context) error {
	response := models.NewResponse(true, "Welcome to Echo Web App!", nil)
	return c.JSON(http.StatusOK, response)
}

// HealthCheck handles the health check endpoint
func HealthCheck(c echo.Context) error {
	response := models.NewResponse(true, "Service is healthy", map[string]string{
		"status": "healthy",
	})
	return c.JSON(http.StatusOK, response)
}

// SetupRoutes configures all the routes for the application
func SetupRoutes(e *echo.Echo) {
	// Public routes
	e.GET("/", Hello)
	e.GET("/health", HealthCheck)

	// API group
	api := e.Group("/api")
	api.GET("/status", func(c echo.Context) error {
		response := models.NewResponse(true, "Application status", map[string]string{
			"version": "1.0.0",
			"name":    "Echo Web App",
		})
		return c.JSON(http.StatusOK, response)
	})

	// Users example endpoint
	api.GET("/users/example", func(c echo.Context) error {
		exampleUser := models.User{
			Username: "testuser",
			Email:    "test@example.com",
		}
		response := models.NewResponse(true, "Example user retrieved", exampleUser)
		return c.JSON(http.StatusOK, response)
	})
}