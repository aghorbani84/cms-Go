package handlers

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"ginwebapp/config"
	"ginwebapp/models"
	"gorm.io/gorm"
)

type AuthHandler struct {
	DB     *gorm.DB
	Config *config.AppConfig
}

func (h *AuthHandler) RequireRole(allowedRoles []string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user := c.Get("user").(*jwt.Token)
			claims := user.Claims.(jwt.MapClaims)
			role, ok := claims["role"].(string)
			
			if !ok || !contains(allowedRoles, role) {
				return c.JSON(http.StatusForbidden, models.NewResponse(false, "Insufficient permissions", nil))
			}
			return next(c)
		}
	}
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func (h *AuthHandler) Login(c echo.Context) error {
	type loginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var req loginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, models.NewResponse(false, "Invalid request", nil))
	}

	var user models.User
	if result := h.DB.Where("email = ?", req.Email).First(&user); result.Error != nil {
		return c.JSON(http.StatusUnauthorized, models.NewResponse(false, "Invalid credentials", nil))
	}

	// In real implementation, verify password hash here
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
		"role": user.Role,
	})

	tokenString, err := token.SignedString([]byte(h.Config.JWTSecret))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.NewResponse(false, "Failed to generate token", nil))
	}

	return c.JSON(http.StatusOK, models.NewResponse(true, "Login successful", map[string]string{
		"token": tokenString,
	}))
}

func (h *AuthHandler) Register(c echo.Context) error {
	// Registration logic here
	return c.JSON(http.StatusNotImplemented, models.NewResponse(false, "Not implemented", nil))
}