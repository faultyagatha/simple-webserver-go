package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

//WriteToConsole logs the page to the console
func WriteToConsole(next http.Handler) http.Handler {
	//cast anonymous func into HandlerFunc type
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("See the page")
		//go to the next middleware
		next.ServeHTTP(w, r)
	})
}

//NoSurf prevents Cross-Site Request Forgery Attacks
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

//SessionLoad loads and saves the cookie sessions on every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
