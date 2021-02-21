package handlers

import (
	"net/http"

	"github.com/faultyagatha/simple-webserver-go/pkg/render"
)

//Home page of the app
func Home(w http.ResponseWriter, r *http.Request) {
	render.Render(w, "home.page.tmpl")
}

//About page of the app
func About(w http.ResponseWriter, r *http.Request) {
	render.Render(w, "about.page.tmpl")
}
