package main

import (
	"fmt"
	"net/http"
)

// spawn a basic http server using net/http

func main() {
	http.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello world! you've requsted: %s\n", r.URL.Path)
	})

	http.ListenAndServe(":8080", nil)
}
