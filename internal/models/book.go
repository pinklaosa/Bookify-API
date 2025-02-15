package models

type Book struct {
	ID            int     `json:"id" gorm:"primaryKey;autoIncrement"`
	Title         string  `json:"title"`
	Author        string  `json:"author"`
	Description   string  `json:"description"`
	CategoryID    int     `json:"category_id"`
	YearPublished int     `json:"year_published"`
	Rating        float32 `json:"rating"`
}
