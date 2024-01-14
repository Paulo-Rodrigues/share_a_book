package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render(w, "homepage.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	render(w, "about.page.tmpl")
}

func render(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)

	err := parsedTemplate.Execute(w, nil)

	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println("Server starting at port 8000")

	http.ListenAndServe(":8000", nil)
}
