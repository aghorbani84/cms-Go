package models

import (
	"time"
	"gorm.io/gorm"
)

type Content struct {
	gorm.Model
	Title       string `gorm:"not null"`
	Slug        string `gorm:"uniqueIndex;not null"`
	Body        string `gorm:"type:text;not null"`
	Status      string `gorm:"default:'draft';not null"` // draft, published, archived
	PublishDate *time.Time
	AuthorID    uint
	Author      User `gorm:"foreignKey:AuthorID"`
	Categories  []Category `gorm:"many2many:content_categories;"`
}

type Category struct {
	gorm.Model
	Name        string `gorm:"uniqueIndex;not null"`
	Description string
}