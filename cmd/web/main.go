package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/faultyagatha/simple-webserver-go/pkg/config"
	"github.com/faultyagatha/simple-webserver-go/pkg/handlers"
	"github.com/faultyagatha/simple-webserver-go/pkg/render"
)

//visible to the main package including middleware and routes
const port string = ":8080"

var app config.AppConfig

var session *scs.SessionManager

func main() {

	//set to true for production
	app.InProduction = false
	session = scs.New()
	//session lasts 24 hours
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	tc, err := render.MakeTmplCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}
	app.TmplCache = tc
	app.UseCache = false
	app.Session = session

	//create a new repository
	repo := handlers.NewRepo(&app)
	//pass the repository back to the handlers
	handlers.NewHandlers(repo)

	render.ConfigTemplate(&app)

	fmt.Println(fmt.Sprintf("Listening on port %s", port))
	// _ = http.ListenAndServe(port, nil)

	serve := &http.Server{
		Addr:    port,
		Handler: routes(&app),
	}
	err = serve.ListenAndServe()
	log.Fatal(err)
}
