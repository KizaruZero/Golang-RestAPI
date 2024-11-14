package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"golang-pustaka-api/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)



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

type BookHandler struct {
	bookService models.Service
}

// Studi Kasus
func NewBookHandler(bookService models.Service) *BookHandler {
	return &BookHandler{
		bookService,
	}
}

// Agar bisa di panggil di main.go, maka harus di export (Huruf depannya huruf besar) wajib nama fungsi huruf depannya besar
func (handler *BookHandler) RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name":    "Ardya Pusaka",
		"message": "Hello, adas!",
	})
}

func (handler *BookHandler) GetBooksHandler(c *gin.Context) {
	books, err := handler.bookService.GetBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, books)
}

func (handler *BookHandler) BookDetailHandler(c *gin.Context) {
	id := c.Param("id")
	bookID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, err := handler.bookService.GetBookByID(bookID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, book)
}

func (handler *BookHandler) BookPostHandler(c *gin.Context) {
	var BookRequest models.BookRequest
	err := c.ShouldBindJSON(&BookRequest)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			message := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, message)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})
		return
	}

	book, err := handler.bookService.CreateBook(BookRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}

func (handler *BookHandler) BookUpdateHandler(c *gin.Context) {
	id := c.Param("id")
	bookID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var book models.Book
	err = c.ShouldBindJSON(&book)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			message := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, message)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})
		return
	}

	updatedBook, err := handler.bookService.UpdateBook(bookID, book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedBook)
}



