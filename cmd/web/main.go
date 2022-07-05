package main

import (
	"fmt"
	"hotel_reservation/pkg/handlers"
	"net/http"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("🚀 Server running on http://localhost%s\n", portNumber)
	http.ListenAndServe(":8080", nil)
}
