# Echo Web Application

A simple web application built with the Echo framework in Go.

## Project Structure

```
ginwebapp/
├── config/         # Application configuration
├── handlers/       # HTTP request handlers
├── middleware/     # Custom middleware functions
├── models/         # Data models and structures
├── static/         # Static files (HTML, CSS, JS)
├── main.go         # Main application entry point
├── go.mod          # Go module definition
└── README.md       # This file
```

## Features

- RESTful API endpoints
- Middleware for logging, recovery, and CORS
- Standardized JSON response format
- Static file serving
- Health check endpoint
- Configuration management

## Getting Started

### Prerequisites

- Go 1.21 or higher

### Installation

1. Clone the repository
2. Navigate to the project directory
3. Install dependencies:

```bash
go mod tidy
```

### Running the Application

```bash
go run main.go
```

The server will start on port 8080. You can access the following endpoints:

- `http://localhost:8080/` - Welcome message
- `http://localhost:8080/health` - Health check endpoint

## Adding New Routes

To add new routes, modify the `main.go` file and add your handler functions.

```go
// Example of adding a new route
e.GET("/api/data", getDataHandler)

// Handler function
func getDataHandler(c echo.Context) error {
    // Your logic here
    return c.JSON(http.StatusOK, yourData)
}
```