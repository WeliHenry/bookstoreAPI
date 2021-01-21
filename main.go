package main

import (
	"encoding/json"
	"fmt"
	"github.com/Weli-Henry/src/github.com/gorilla/mux"
	"math/rand"
	"net/http"
	"strconv"
)

// book struct declaration

type Book struct {
	ID     string `json:"id"`
	Isbn   string `json:"isbn"`
	Title  string `json:"title"`
	Author *Author
}
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var books []Book

func main() {

	fmt.Println("application started successfully")
	r := mux.NewRouter()

	books = append(books, Book{
		ID:    "1",
		Isbn:  "45545",
		Title: "book one",
		Author: &Author{
			Firstname: "Matthew",
			Lastname:  "Simpson",
		},
	})
	books = append(books, Book{
		ID:    "2",
		Isbn:  "78545",
		Title: "book two",
		Author: &Author{
			Firstname: "johnny",
			Lastname:  "holmes",
		},
	})
	//handlers declaration
	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/books", createBook).Methods("POST")
	r.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")
	http.ListenAndServe(":9001", r)
}

//fetching all books from the memory slice
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(books)

}

//fetching a single book from the local memory slice
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)

	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

//adding a single book to the local memory slice
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(10000000))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

//updating a single book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)

	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			w.Header().Set("content-type", "application/json")
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
		json.NewEncoder(w).Encode(books)
	}
}

//deleting a single book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)

	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(books)
	}
}
