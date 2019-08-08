package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/albertowusuasare/customer-app/app"
	"github.com/albertowusuasare/customer-app/app/google"
)

func main() {
	ctx := context.Background()
	firestoreClient, err := firestore.NewClient(ctx, "onua-246719")
	if err != nil {
		log.Fatal(err)
	}
	googleApp := google.App(ctx, firestoreClient)
	port := ":5090"
	log.Println(fmt.Sprintf("Starting server on port%s", port))
	log.Fatal(http.ListenAndServe(port, app.Handler(googleApp)))
}
