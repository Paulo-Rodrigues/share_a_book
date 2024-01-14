package main

import (
	"fmt"
	"net/http"
	"share_a_book/pkg/handlers"
)

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println("Server starting at port 8000")

	http.ListenAndServe(":8000", nil)
}
