package main

import (
	"customer-app/cmd/customer-svc/handler"
	"customer-app/cmd/customer-svc/pkg"
	"fmt"
	"log"
	"net/http"
)

func main() {
	app := pkg.InmemApp()
	port := ":5090"
	log.Println(fmt.Sprintf("Starting server on port%s", port))
	log.Fatal(http.ListenAndServe(port, handler.Handle(app)))
}
