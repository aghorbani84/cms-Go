package main

import (
	"fmt"
	"ginwebapp/config"
	"ginwebapp/handlers"
	"ginwebapp/middleware"

	"github.com/labstack/echo/v4"
)

func main() {
	// Load configuration
	appConfig := config.GetConfig()

	// Create a new Echo instance
	e := echo.New()

	// Setup middleware from middleware package
	middleware.SetupMiddleware(e)

	// Setup routes from handlers package
	handlers.SetupRoutes(e)

	// Serve static files
	e.Static("/static", "static")
	
	// Serve index.html at root path as an alternative to the API response
	e.File("/web", "static/index.html")

	// Start server
	serverAddr := fmt.Sprintf(":%s", appConfig.Port)
	fmt.Printf("Server starting on port %s in %s mode\n", 
		appConfig.Port, appConfig.Environment)
	fmt.Printf("Visit http://localhost:%s/web to see the web interface\n", appConfig.Port)
	e.Logger.Fatal(e.Start(serverAddr))
}