package handlers

import (
	"net/http"
	"share_a_book/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.Render(w, "homepage.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.Render(w, "about.page.tmpl")
}
