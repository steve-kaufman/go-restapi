package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"

	"github.com/gorilla/mux"
)

// Book Struct (Model)
type Book struct {
	gorm.Model
	Isbn     string `json:"isbn"`
	Title    string `json:"title"`
	AuthorID int    `json:"authorId"`
	Author   Author `json:"author"`
}

// Author Struct
type Author struct {
	gorm.Model
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Error Struct
type Error struct {
	Message string
}

// Init database
var db *gorm.DB

// Get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var books []Book
	db.Find(&books)
	json.NewEncoder(w).Encode(books)
}

// Get one book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var id string = mux.Vars(r)["id"]

	var book Book

	db.First(&book, id)

	// check if book was found
	if (book != Book{}) {
		json.NewEncoder(w).Encode(book)
		return
	}

	w.WriteHeader(404)
	json.NewEncoder(w).Encode(Error{Message: "Book with id " + id + " not found"})
}

// Create a book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)

	db.Create(&book)

	w.WriteHeader(201)
	json.NewEncoder(w).Encode(book)
}

// Update a book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var id string = mux.Vars(r)["id"]

	var book Book

	db.First(&book, id)

	body, _ := ioutil.ReadAll(r.Body)

	if (book != Book{}) {
		json.Unmarshal(body, &book)

		db.Save(&book)

		json.NewEncoder(w).Encode(book)
		return
	}

	w.WriteHeader(404)
	json.NewEncoder(w).Encode(Error{Message: "Book not found"})
}

// Delete a book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var id string = mux.Vars(r)["id"]

	var book Book

	db.First(&book, id)

	if (book != Book{}) {
		db.Delete(&book)
		json.NewEncoder(w).Encode(book)
		return
	}

	w.WriteHeader(404)
	json.NewEncoder(w).Encode(Error{Message: "Book not found"})
}

func main() {
	// Init database
	db = CreateDatabase()

	// Init Router
	router := mux.NewRouter()

	// Route handlers
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PATCH")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
