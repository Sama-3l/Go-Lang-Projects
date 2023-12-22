package models

import (
	"sqlBookManagement/pkg/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
	Tags        []Tags `gorm:"many2many:book_tags;" json:"tags"`
}

type Tags struct {
	gorm.Model
	Tag_Name string `gorm:"unique" json:"tag-name"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{}, &Tags{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	for i := range Books {
		db.Model(&Books[i]).Association("Tags").Find(&Books[i].Tags)
	}
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID = ?", Id).Preload("Tags").Find(&getBook)
	return &getBook, db
}

func DeleteBook(Id int64) Book {
	var book Book
	db.Preload("Tags").First(&book, Id)
	db.Delete(&book)
	for _, tag := range book.Tags {
		db.Model(&book).Association("Tags").Delete(&tag)
		db.Delete(&tag)
	}
	return book
}
