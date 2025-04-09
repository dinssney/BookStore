package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Publisher   User   `json:"publisher"`
	PublisherID uint   `json:"publisher_id"`
}
