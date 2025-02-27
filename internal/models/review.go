package models

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	UserID  int    `json:"user_id"`
	BookID  int    `json:"book_id"`
	Rating  int    `json:"rating"`
	Comment string `json:"comment"`
}
