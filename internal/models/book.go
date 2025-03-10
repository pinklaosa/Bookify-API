package models

import "gorm.io/gorm"

type Book struct {
	// ID            int     `json:"id" gorm:"primaryKey;autoIncrement"`
	gorm.Model
	Title         string  `json:"title"`
	Author        string  `json:"author"`
	Description   string  `json:"description"`
	CategoryID    int     `json:"category_id"`
	YearPublished int     `json:"year_published"`
	Rating        float32 `json:"rating"`
}


type GetBookParams struct {
	Title string
	Category string
	Rating string
}