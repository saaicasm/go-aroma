package handler

import (
	"github/iamlego/go-web/pkg/render"
	"net/http"
)

// Function handler for Home
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")

}

// Function handler for About
func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
