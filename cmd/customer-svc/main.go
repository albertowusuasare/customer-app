package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/albertowusuasare/customer-app/internal/app"
	"github.com/albertowusuasare/customer-app/internal/app/inmem"
)

func main() {
	inmemApp := inmem.App()
	port := ":5090"
	log.Println(fmt.Sprintf("Starting server on port%s", port))
	log.Fatal(http.ListenAndServe(port, app.Handler(inmemApp)))
}
