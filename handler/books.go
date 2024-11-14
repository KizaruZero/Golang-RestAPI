package handler

import (
	"net/http"
	"strconv"

	"golang-pustaka-api/models"

	"github.com/gin-gonic/gin"
)

// Agar bisa di panggil di main.go, maka harus di export (Huruf depannya huruf besar) wajib nama fungsi huruf depannya besar
func RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name":    "Ardya Pusaka",
		"message": "Hello, adas!",
	})
}

func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
		"status":  http.StatusOK,
	})
}

func QueryHandler(c *gin.Context) {
	id := c.Query("id")
	page := c.DefaultQuery("page", "1")
	c.JSON(http.StatusOK, gin.H{
		"id":   id,
		"page": page,
	})
}



var books = []models.Book{
	{ID: 1, Title: "The Go Programming Language", Author: "Alan A. A. Donovan", Price: 500000},
	{ID: 1, Title: "Introducing Go", Author: "Caleb Doxsey", Price: 300000},
}

func GetBooksHandler(c *gin.Context) {
	c.JSON(http.StatusOK, books)
}

func BookDetailHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32) // Konversi id dari string ke uint64
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}

	for _, book := range books {
		if book.ID == uint(id) { // Bandingkan dengan book.ID setelah konversi ke uint
			c.JSON(http.StatusOK, book)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
}

func BookPostHandler(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	books = append(books, book)
	c.JSON(http.StatusCreated, book)
}

// contoh youtube :
func PostBookHandler (c *gin.Context) {
	var book models.Book
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
