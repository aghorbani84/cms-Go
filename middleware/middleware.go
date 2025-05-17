package middleware

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// SetupMiddleware configures all middleware for the application
func SetupMiddleware(e *echo.Echo) {
	// Basic middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	
	// CORS middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))
	
	// Request ID middleware
	e.Use(middleware.RequestID())
	
	// Custom middleware for request timing
	e.Use(RequestTimer)
}

// RequestTimer is a middleware that times requests
func RequestTimer(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()
		
		err := next(c)
		
		// Calculate request time
		stop := time.Now()
		duration := stop.Sub(start)
		
		// Add header with request duration
		c.Response().Header().Set("X-Response-Time", duration.String())
		
		return err
	}
}