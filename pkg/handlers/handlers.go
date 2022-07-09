package handlers

import (
	"hotel_reservation/pkg/config"
	"hotel_reservation/pkg/render"
	"net/http"
)

//TemplateData holds data sent from handlers to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	flash     string
	warning   string
	error     string
}

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// Repo the repository used by the handlers
var Repo *Repository

// InitRepo creates a new repository
func InitRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (re *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "home.page.tmpl", &TemplateData{})
}

// About is the about page handler
func (re *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "about.page.tmpl", &TemplateData{})
}
g