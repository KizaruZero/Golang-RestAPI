package models

// Book struct
type BookRequest struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
	Price  int    `json:"price" binding:"required"`
}
