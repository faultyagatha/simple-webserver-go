package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/faultyagatha/simple-webserver-go/pkg/config"
	"github.com/faultyagatha/simple-webserver-go/pkg/handlers"
	"github.com/faultyagatha/simple-webserver-go/pkg/render"
)

const port string = ":8080"

func main() {
	var app config.AppConfig
	tc, err := render.MakeTmplCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}
	app.TmplCache = tc
	app.UseCache = false

	//create a new repository
	repo := handlers.NewRepo(&app)
	//pass the repository back to the handlers
	handlers.NewHandlers(repo)

	render.ConfigTemplate(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)
	fmt.Println(fmt.Sprintf("Listening on port %s", port))
	_ = http.ListenAndServe(port, nil)
}
