package main

import (
	"fmt"
	"net/http"
)

// use net/http to spawn a basic http server
// use that server to serve static assets on registered route

func main() {
	http.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, I am alive...")
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":8080", nil)
}
