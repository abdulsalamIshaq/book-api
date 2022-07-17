package main

import (
	"log"

	"github.com/abdulsalamIshaq/book-api/models"
	"github.com/abdulsalamIshaq/book-api/services"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := models.Database()

	if err != nil {
		log.Println(err)
	}

	db.DB()

	router := gin.Default()

	router.GET("/books", services.GetBooks)
	router.GET("/books/:id", services.GetBook)
	router.POST("/books", services.PostBook)
	router.PUT("/books/:id", services.UpdateBook)
	router.DELETE("/books/:id", services.DeleteBook)

	log.Fatal(router.Run(":8000"))
}
