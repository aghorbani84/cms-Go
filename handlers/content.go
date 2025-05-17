package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"ginwebapp/config"
	"ginwebapp/models"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type ContentHandler struct {
	DB     *gorm.DB
	Config *config.AppConfig
}

func (h *ContentHandler) CreateContent(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	
	roleFloat, ok := claims["role"].(float64)
	if !ok || (int(roleFloat) != models.RoleAdmin && int(roleFloat) != models.RoleEditor) {
		return c.JSON(http.StatusForbidden, models.NewResponse(false, "Insufficient permissions", nil))
	}

	var content models.Content
	if err := c.Bind(&content); err != nil {
		return c.JSON(http.StatusBadRequest, models.NewResponse(false, "Invalid request", nil))
	}

	content.AuthorID = uint(claims["sub"].(float64))
	if result := h.DB.Create(&content); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.NewResponse(false, "Failed to create content", nil))
	}

	return c.JSON(http.StatusCreated, models.NewResponse(true, "Content created", content))
}

func (h *ContentHandler) ListContent(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	roleFloat, ok := claims["role"].(float64)
	if !ok {
		return c.JSON(http.StatusForbidden, models.NewResponse(false, "Access denied", nil))
	}
	role := int(roleFloat)

	var contents []models.Content
	query := h.DB.Preload("Categories")

	switch role {
	case models.RoleAdmin:
		// Admins see all content
	case models.RoleEditor:
		query = query.Where("author_id = ?", claims["sub"])
	default:
		return c.JSON(http.StatusForbidden, models.NewResponse(false, "Access denied", nil))
	}

	if result := query.Find(&contents); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.NewResponse(false, "Failed to retrieve content", nil))
	}

	return c.JSON(http.StatusOK, models.NewResponse(true, "Content list", contents))
}

func (h *ContentHandler) UpdateContent(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	roleFloat, ok := claims["role"].(float64)
	if !ok {
		return c.JSON(http.StatusForbidden, models.NewResponse(false, "Insufficient permissions", nil))
	}
	role := int(roleFloat)

	var content models.Content
	if result := h.DB.First(&content, c.Param("id")); result.Error != nil {
		return c.JSON(http.StatusNotFound, models.NewResponse(false, "Content not found", nil))
	}

	if role != models.RoleAdmin && content.AuthorID != uint(claims["sub"].(float64)) {
		return c.JSON(http.StatusForbidden, models.NewResponse(false, "Insufficient permissions", nil))
	}

	if err := c.Bind(&content); err != nil {
		return c.JSON(http.StatusBadRequest, models.NewResponse(false, "Invalid request", nil))
	}

	if result := h.DB.Save(&content); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.NewResponse(false, "Failed to update content", nil))
	}

	return c.JSON(http.StatusOK, models.NewResponse(true, "Content updated", content))
}

func (h *ContentHandler) DeleteContent(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	roleFloat, ok := claims["role"].(float64)
	if !ok {
		return c.JSON(http.StatusForbidden, models.NewResponse(false, "Insufficient permissions", nil))
	}
	role := int(roleFloat)

	var content models.Content
	if result := h.DB.First(&content, c.Param("id")); result.Error != nil {
		return c.JSON(http.StatusNotFound, models.NewResponse(false, "Content not found", nil))
	}

	if role != models.RoleAdmin && content.AuthorID != uint(claims["sub"].(float64)) {
		return c.JSON(http.StatusForbidden, models.NewResponse(false, "Insufficient permissions", nil))
	}

	if result := h.DB.Delete(&content); result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.NewResponse(false, "Failed to delete content", nil))
	}

	return c.JSON(http.StatusOK, models.NewResponse(true, "Content deleted", nil))
}