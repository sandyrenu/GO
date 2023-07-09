package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/your-username/book-service/models"
	"github.com/your-username/book-service/repositories"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	books := repositories.GetAllBooks()
	json.NewEncoder(w).Encode(books)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookID := params["id"]
	book := repositories.GetBookByID(bookID)
	json.NewEncoder(w).Encode(book)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	createdBook := repositories.CreateBook(book)
	json.NewEncoder(w).Encode(createdBook)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookID := params["id"]

	var book models.Book
	_ = json.NewDecoder(r.Body).Decode(&book)

	updatedBook := repositories.UpdateBook(bookID, book)
	json.NewEncoder(w).Encode(updatedBook)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookID := params["id"]
	repositories.DeleteBook(bookID)
}
