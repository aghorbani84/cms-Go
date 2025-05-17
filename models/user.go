package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string `gorm:"uniqueIndex;not null"`
	Email        string `gorm:"uniqueIndex;not null"`
	PasswordHash string `gorm:"not null"`
	Role         int    `gorm:"not null;default:3"` // 1=Admin, 2=Editor, 3=Viewer
	LastLogin    *time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

const (
	RoleAdmin = iota + 1
	RoleEditor
	RoleViewer
)