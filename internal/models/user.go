package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	Role     Role   `json:"role"`
	Books    []Book `gorm:"many2many:user_books;"`
}
