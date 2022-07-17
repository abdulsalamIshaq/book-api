package services

import (
	"log"
	"net/http"

	"github.com/abdulsalamIshaq/book-api/models"
	"github.com/gin-gonic/gin"
)

type NewBookStruct struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
	ISBN   string `json:"isbn" binding:"required"`
}

type UpdateBookStruct struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	ISBN   string `json:"isbn"`
}

func GetBooks(c *gin.Context) {

	var books []models.Book

	db, err := models.Database()

	if err != nil {
		log.Println(err)
	}

	if err := db.Find(&books).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, books)

}

func GetBook(c *gin.Context) {
	var book models.Book

	db, err := models.Database()
	if err != nil {
		log.Println(err)
	}

	if err := db.Where("id=?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	c.JSON(http.StatusOK, book)
}

func PostBook(c *gin.Context) {
	var book NewBookStruct

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newBook := models.Book{Title: book.Title, Author: book.Author, ISBN: book.ISBN}

	db, err := models.Database()
	if err != nil {
		log.Println(err)
	}

	if err := db.Create(&newBook).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, newBook)
}

func UpdateBook(c *gin.Context) {
	var book models.Book

	db, err := models.Database()

	if err != nil {
		log.Println(err)
	}

	if err := db.Where("id=?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	var updateBook UpdateBookStruct

	if err := c.ShouldBindJSON(&updateBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Model(&book).Updates(models.Book{Title: updateBook.Title, Author: updateBook.Author, ISBN: updateBook.ISBN}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, book)
}

func DeleteBook(c *gin.Context) {

	var book models.Book

	db, err := models.Database()

	if err != nil {
		log.Println(err)
	}

	if err := db.Where("id=?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	if err := db.Delete(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "book deleted successfully"})

}
