package handlers

import (
	"net/http"

	"github.com/faultyagatha/simple-webserver-go/pkg/config"
	"github.com/faultyagatha/simple-webserver-go/pkg/render"
)

// Repo is the repository used by the handlers
var Repo *Repository

// Repository is the repository type that holds a pointer to the config
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.Render(w, "home.page.tmpl")
}

// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.Render(w, "about.page.tmpl")
}
