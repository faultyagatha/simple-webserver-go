package main

import (
	"fmt"
	"net/http"

	"github.com/faultyagatha/simple-webserver-go/pkg/handlers"
)

const port string = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	fmt.Println(fmt.Sprintf("Listening on port %s", port))
	_ = http.ListenAndServe(port, nil)
}
