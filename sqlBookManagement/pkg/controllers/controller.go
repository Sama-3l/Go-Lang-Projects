package controllers

import (
	"encoding/json"
	"net/http"
	"sqlBookManagement/pkg/models"
)

var NewBook models.Book

func CreateBook(w http.ResponseWriter, req *http.Request) {

}

func GetBook(w http.ResponseWriter, req *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, req *http.Request) {

}

func UpdateBook(w http.ResponseWriter, req *http.Request) {

}

func DeleteBook(w http.ResponseWriter, req *http.Request) {

}
