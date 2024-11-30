package handler

import (
	"github/iamlego/go-web/pkg/config"
	"github/iamlego/go-web/pkg/render"
	"net/http"
)

// Repository pattern

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandler(r *Repository) {
	Repo = r
}

// Function handler for Home
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")

}

// Function handler for About
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}
