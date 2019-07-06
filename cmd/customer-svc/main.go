package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/albertowusuasare/customer-app/cmd/customer-svc/handler"
	"github.com/albertowusuasare/customer-app/cmd/customer-svc/pkg"
)

func main() {
	app := pkg.InmemApp()
	port := ":5090"
	log.Println(fmt.Sprintf("Starting server on port%s", port))
	log.Fatal(http.ListenAndServe(port, handler.Handle(app)))
}
