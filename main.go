package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Home Page")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the About Page")
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("🚀 Server running on http://localhost%s\n", portNumber)
	http.ListenAndServe(":8080", nil)
}
