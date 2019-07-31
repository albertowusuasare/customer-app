package google

import (
	"context"
	"log"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/albertowusuasare/customer-app/internal/api"
	"github.com/albertowusuasare/customer-app/internal/app"
)

//Handle is the entry point for the cloud function
func Handle(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	firestoreClient, err := firestore.NewClient(ctx, "onua-246719")
	if err != nil {
		log.Fatal(err)
	}
	app := app.GoogleApp(ctx, firestoreClient)
	handler := api.Handler(app)
	handler.ServeHTTP(w, r)
}
