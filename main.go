package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Book struct {
	ID     int    `json:id`
	Title  string `json:title`
	Author string `json:author`
	Year   string `json:year`
}

var books []Book

func main() {

	router := mux.NewRouter()

	books = append(books, Book{ID: 1, Title: "Golang pointers", Author: "Mr. Golang", Year: "2010"},
		Book{ID: 2, Title: "Goroutines", Author: "Mr.Gourity", Year: "1989"},
		Book{ID: 3, Title: "Go routers", Author: "Mr.Balddy", Year: "1923"},
		Book{ID: 4, Title: "Go pickme up", Author: "Mr.GBad", Year: "1921"},
		Book{ID: 5, Title: "Go down on it", Author: "Mr.Good", Year: "1911"})

	router.HandleFunc("/books", getBook).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))

}

func getBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Get one book")
}

func addBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Add one book")
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Updates a book")
}

func removeBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Removes a book")
}
