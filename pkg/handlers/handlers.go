package handlers

import (
	"hotel_reservation/pkg/config"
	"hotel_reservation/pkg/render"
	"net/http"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

//NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (re *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "home.page.tmpl")
}

// About is the about page handler
func (re *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "about.page.tmpl")
}
