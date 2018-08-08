package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Book Struct (Model)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author Struct
type Author Struct {
	Firstname 	string 	`json:"firstname`
	Lastname 	string 	`json:"lastname`
}

// Get All Books
func getBooks(w http.ResponseWriter, r *http.Request)
{

}

// Get Single Book
func getBook(w http.ResponseWriter, r *http.Request)
{
	
}

// Create a New Book
func createBook(w http.ResponseWriter, r *http.Request)
{
	
}

// Update an Existing Book
func updateBook(w http.ResponseWriter, r *http.Request)
{
	
}

// Delete an Existing Book
func deleteBook(w http.ResponseWriter, r *http.Request)
{
	
}
func main() {
	// Init Router
	r := mux.NewRouter()

	// Route Handlers / Endpoints
	r.HandleFunc("api/books", getBooks).Methods(GET)
	r.HandleFunc("api/books/{id}", getBook).Methods(GET)
	r.HandleFunc("api/books", createBook).Methods(POST)
	r.HandleFunc("api/books/{id}", updateBook).Methods(PUT)
	r.HandleFunc("api/books/{id}", deleteBook).Methods(DELETE)

	log.Fatal(http.ListenAndServe(":8000", r))
}
