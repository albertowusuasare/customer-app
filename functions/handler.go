package functions

import (
	"context"
	"log"
	"net/http"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/pubsub"
	"github.com/albertowusuasare/customer-app/app"
	"github.com/albertowusuasare/customer-app/app/google"
)

//Handle is the entry point for the cloud function
func Handle(w http.ResponseWriter, r *http.Request) {
	projectID := retrieveProjectID()
	ctx := context.Background()
	firestoreClient, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatal(err)
	}
	pubsubClient, pubsubErr := pubsub.NewClient(ctx, projectID)
	if pubsubErr != nil {
		log.Fatal(pubsubErr)
	}
	googleApp := google.App(ctx, firestoreClient, pubsubClient)
	handler := app.Handler(googleApp)
	handler.ServeHTTP(w, r)
}

func retrieveProjectID() string {
	return "onua-246719"
}
