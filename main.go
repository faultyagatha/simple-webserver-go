package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const port string = ":8080"

//Home page of the app
func Home(w http.ResponseWriter, r *http.Request) {
	render(w, "home.tmpl")
}

//About page of the app
func About(w http.ResponseWriter, r *http.Request) {
	render(w, "about.tmpl")
}

//function that renders golang templates
func render(w http.ResponseWriter, tempate string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tempate)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template: ", err)
	}
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	fmt.Println(fmt.Sprintf("Listening on port %s", port))
	_ = http.ListenAndServe(port, nil)
}
