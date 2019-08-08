package functions

import (
	"context"
	"log"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/albertowusuasare/customer-app/internal/app"
	"github.com/albertowusuasare/customer-app/internal/app/google"
)

//Handle is the entry point for the cloud function
func Handle(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	firestoreClient, err := firestore.NewClient(ctx, "onua-246719")
	if err != nil {
		log.Fatal(err)
	}
	googleApp := google.App(ctx, firestoreClient)
	handler := app.Handler(googleApp)
	handler.ServeHTTP(w, r)
}
