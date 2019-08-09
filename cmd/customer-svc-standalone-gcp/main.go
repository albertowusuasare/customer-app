package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/pubsub"
	"github.com/albertowusuasare/customer-app/app"
	"github.com/albertowusuasare/customer-app/app/google"
)

func main() {
	projectID := retrieveProjectID()
	ctx := context.Background()
	firestoreClient, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatal(err)
	}
	pubsubClient, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Fatal(err)
	}
	googleApp := google.App(ctx, firestoreClient, pubsubClient)
	port := ":5090"
	log.Println(fmt.Sprintf("Starting server on port%s", port))
	log.Fatal(http.ListenAndServe(port, app.Handler(googleApp)))
}

func retrieveProjectID() string {
	return "onua-246719"
}
