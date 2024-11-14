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
	bookService := models.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)
	


	// This is a default router
	router := gin.Default()

	// Api Versioning
	// Jadi studi kasusnya misal ada 2 versi API, versi 1 dan versi 2 yang berbeda struktur API nya maka bisa menggunakan cara ini agar tidak bentrok
	v1 := router.Group("/v1")
	v1.GET("/", bookHandler.RootHandler) 
	v1.GET("/books", bookHandler.GetBooksHandler)
	v1.POST("/book", bookHandler.BookPostHandler)
	v1.GET("/book/:id", bookHandler.BookDetailHandler) 
	v1.PUT("/book/:id", bookHandler.BookUpdateHandler)

	
	router.Run(":8080")
}

