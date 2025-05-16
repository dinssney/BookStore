package models

import "gorm.io/gorm"

// test

type Book struct {
	gorm.Model
	Title         string `json:"title" binding:"required"`
	Author        string `json:"author" binding:"required"`
	Description   string `json:"description"`
	Publisher     User   `json:"publisher" gorm:"foreignKey:PublisherID"`
	PublisherID   uint   `json:"publisher_id"`
	CoverImageUrl string `json:"cover_image_url" binding:"required"`
}
