package render

import (
	"fmt"
	"html/template"
	"net/http"
)

func Render(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")

	err := parsedTemplate.Execute(w, nil)

	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}
}
