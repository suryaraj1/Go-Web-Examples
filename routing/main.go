package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	fmt.Fprintf(w, "Creating a book: %s\n", title)
}

func ReadBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	fmt.Fprintf(w, "Reading the book: %s\n", title)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	})

	r.HandleFunc("/books/{title}", CreateBook).Methods("POST")
	r.HandleFunc("/books/{title}", ReadBook).Methods("GET")

	http.ListenAndServe(":8080", r)
}
