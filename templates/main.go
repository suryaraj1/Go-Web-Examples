package main

import (
	"fmt"
	"net/http"
	"text/template"
)

// have the layout in an html file using GO templates
// parse the template from the file on disk
// register a request handler and make it execute the template

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func main() {
	// fmt.Println("hello")

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Server is doing good..")
	})

	tmpl := template.Must(template.ParseFiles("layout.html"))
	http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			PageTitle: "My Demo Todo List",
			Todos: []Todo{
				{Title: "Task 1", Done: false},
				{Title: "Task 2", Done: false},
				{Title: "Task 3", Done: true},
			},
		}

		tmpl.Execute(w, data)
	})

	http.ListenAndServe(":8080", nil)
}
