package models

type Book struct {
	ID           int     `json:"id"`
	Title        string  `json:"title"`
	Author       string  `json:"author"`
	Description  string  `json:"description"`
	CategetoryID int     `json:"categetory_id"`
	YearPubliced int     `json:"year_publiced"`
	Rating       float32 `json:"rating"`
}
