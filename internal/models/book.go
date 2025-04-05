package models

import "gorm.io/gorm"

// test

type Book struct {
	gorm.Model
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
}
