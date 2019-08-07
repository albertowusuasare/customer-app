package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/albertowusuasare/customer-app/internal/app"
)

func main() {
	inmemApp := app.Inmem()
	port := ":5090"
	log.Println(fmt.Sprintf("Starting server on port%s", port))
	log.Fatal(http.ListenAndServe(port, app.Handler(inmemApp)))
}
