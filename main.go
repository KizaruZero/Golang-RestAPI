package main

import (
	"golang-pustaka-api/handler"
	"golang-pustaka-api/models"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// Database Connection
	dsn := "root:@tcp(127.0.0.1:3306)/pustaka-api?charset=utf8mb4&parseTime=True&loc=Local"
  	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database!")
	}
	db.AutoMigrate(&models.Book{}) // wajib pointer
	
	// Repository
	bookRepository := models.NewRepository(db)

	// Implementasi Repo 
	books, err := bookRepository.GetBooks()
	if err != nil {
		panic("Failed to get books!")
	}

	for _, book := range books {
		println(book.Title)
	}

	// Create Book
	books2 := models.Book{
		Title:  "Belajar Golang",
		Author: "Ardya Pusaka",
		Price:  100000,
	}

	bookRepository.CreateBook(books2)
	
	


	 

	
	




	// This is a default router
	router := gin.Default()

	// Api Versioning
	// Jadi studi kasusnya misal ada 2 versi API, versi 1 dan versi 2 yang berbeda struktur API nya maka bisa menggunakan cara ini agar tidak bentrok
	v1 := router.Group("/v1")
	v1.GET("/", handler.RootHandler) 
	v1.GET("/ping", handler.PingHandler)
	v1.GET("/books", handler.GetBooksHandler)
	v1.POST("/book", handler.BookPostHandler)
	v1.GET("/book/:id", handler.BookDetailHandler) 
	v1.GET("/query", handler.QueryHandler)

	
	router.Run(":8080")
}

