package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/HlufD/books_api/models"
	"github.com/HlufD/books_api/utils"
	"github.com/gorilla/mux"
)

func CreateBook(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	book := models.Book{}
	utils.ParseJsonBody(req, &book)
	b := models.ScreateBooks(book)
	result, _ := json.Marshal(b)
	res.WriteHeader(http.StatusOK)
	res.Write(result)
}

func GetBooks(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	books, _ := models.SgetAllBooks()
	result, _ := json.Marshal(books)
	res.WriteHeader(http.StatusOK)
	res.Write(result)

}
func GeteBook(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	bookId := params["id"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parssin id")
	}
	book := models.SgetBook(id)
	result, _ := json.Marshal(book)
	res.WriteHeader(http.StatusOK)
	res.Write(result)

}
func UpdateBook(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	book := models.Book{}
	utils.ParseJsonBody(req, &book)
	params := mux.Vars(req)
	bookId := params["id"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parssin id")
	}
	updatedBook, _ := models.SupdateBook(id, &book)
	result, _ := json.Marshal(updatedBook)
	res.WriteHeader(http.StatusOK)
	res.Write(result)
}

func DeleteBook(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)
	bookId := params["id"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parssin id")
	}
	book := models.SdeleteBook(id)
	result, _ := json.Marshal(book)
	res.WriteHeader(http.StatusOK)
	res.Write(result)
}
