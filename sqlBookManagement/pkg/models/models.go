package models

import (
	"fmt"
	"sqlBookManagement/pkg/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
	Tags        []Tags `gorm:"many2many:book_tags;" json:"tags"`
}

type Tags struct {
	gorm.Model
	Tag_Name string `gorm:""json:"tag-name"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{}, &Tags{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	fmt.Println(b)
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
	db.Where("ID=?", Id).Delete(book)
	return book
}
