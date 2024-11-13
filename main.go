package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// This is a comment
	router := gin.Default()
	router.GET("/", rootHandler) 
	router.GET("/ping", pingHandler)
	router.GET("/books", getBooksHandler)
	router.POST("/books", bookPostHandler)
	router.POST("/book", postBookHandler)
	router.GET("/book/:id", bookDetailHandler) 
	router.GET("/query", queryHandler)

	router.Run(":8080")
}

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name":    "Ardya Pusaka",
		"message": "Hello, adas!",
	})
}

func pingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
		"status":  http.StatusOK,
	})
}

func queryHandler(c *gin.Context) {
	id := c.Query("id")
	page := c.DefaultQuery("page", "1")
	c.JSON(http.StatusOK, gin.H{
		"id":   id,
		"page": page,
	})
}

// Struct untuk data buku
type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Price  int    `json:"price"`
}

var books = []Book{
	{ID: "1", Title: "The Go Programming Language", Author: "Alan A. A. Donovan", Price: 500000},
	{ID: "2", Title: "Introducing Go", Author: "Caleb Doxsey", Price: 300000},
}

func getBooksHandler(c *gin.Context) {
	c.JSON(http.StatusOK, books)
}

func bookDetailHandler(c *gin.Context) {
	id := c.Param("id")
	for _, book := range books {
		if book.ID == id {
			c.JSON(http.StatusOK, book)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
}

func bookPostHandler(c *gin.Context) {
	var book Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	books = append(books, book)
	c.JSON(http.StatusCreated, book)
}

// contoh youtube :
func postBookHandler (c *gin.Context) {
	var book Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200,gin.H{
		"title" : book.Title,
		"author" : book.Author,
		"price" : book.Price,
	})
}
