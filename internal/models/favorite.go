package models

import "gorm.io/gorm"

type Favorite struct {
    gorm.Model
    UserID int `json:"user_id" gorm:"not null"`
    BookID int `json:"book_id" gorm:"not null"`
}
