package google

import (
	"net/http"

	"github.com/albertowusuasare/customer-app/internal/api"
	"github.com/albertowusuasare/customer-app/internal/app"
)

//Handle is the entry point for the cloud function
func Handle(w http.ResponseWriter, r *http.Request) {
	app := app.Inmem()
	handler := api.Handler(app)
	handler.ServeHTTP(w, r)
}
