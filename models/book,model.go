package models

import (
	"log"

	"github.com/HlufD/books_api/config"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name        string `gomr:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

var db *gorm.DB

func init() {
	config.ConnectToDb()
	db = config.GetDb()
	if err := db.AutoMigrate(&Book{}); err != nil {
		log.Fatal("Faild to migrate the Database", err)
	}
}

// createting a service methodes that our controllers will call

func ScreateBooks(book Book) Book {
	db.Create(&book)
	return book
}

func SgetAllBooks() ([]Book, error) {
	var books []Book
	if err := db.Unscoped().Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func SgetBook(id int64) Book {
	var book Book
	db.Where("id=?", id).Find(&book)
	return book
}

func SdeleteBook(id int64) Book {
	var book Book
	db.First(&book, id)
	db.Unscoped().Delete(book)
	return book
}

func SupdateBook(id int64, book *Book) (Book, error) {
	var existingBook Book
	if err := db.First(&existingBook, id).Error; err != nil {
		return existingBook, err
	}
	if book.Name != "" {
		existingBook.Name = book.Name
	}
	if book.Author != "" {
		existingBook.Author = book.Author
	}
	if book.Publication != "" {
		existingBook.Publication = book.Publication
	}
	if err := db.Save(&existingBook).Error; err != nil {
		return existingBook, err
	}

	return existingBook, nil
}
