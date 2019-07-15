package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/albertowusuasare/customer-app/internal/api"
	"github.com/albertowusuasare/customer-app/internal/app"
)

func main() {
	app := app.Inmem()
	port := ":5090"
	log.Println(fmt.Sprintf("Starting server on port%s", port))
	log.Fatal(http.ListenAndServe(port, api.Handler(app)))
}
