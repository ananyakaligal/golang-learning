package main

import (
	"fmt"
	"net/http"
)

// hello handler
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Go server!")
}

// form handler
func formHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/form.html")
}

func main() {
	// route for /
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	})

	// route for /hello
	http.HandleFunc("/hello", helloHandler)

	// route for /form
	http.HandleFunc("/form", formHandler)

	// start server on port 8080
	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
