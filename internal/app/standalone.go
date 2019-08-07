package app

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// StandAloneFunc creates a new instance of the stand alone  application
type StandAloneFunc func() StandAlone

// StandAlone encapsulates all the apps as a single standalone app
type StandAlone struct {
	CreateHandler        http.HandlerFunc
	RetrieveOneHandler   http.HandlerFunc
	RetrieveMultiHandler http.HandlerFunc
	UpdateHandler        http.HandlerFunc
	RemoveHandler        http.HandlerFunc
}

// HandlerFunc returns an entry point http handler for the entire application
type HandlerFunc func(app StandAlone) http.Handler

// Handler is the default http handler for the application
func Handler(app StandAlone) http.Handler {
	r := httprouter.New()
	r.Handler(http.MethodPost, "/customers", app.CreateHandler)
	r.Handler(http.MethodGet, "/customers/:customerId", app.RetrieveOneHandler)
	r.Handler(http.MethodGet, "/customers", app.RetrieveMultiHandler)
	r.Handler(http.MethodPut, "/customers/:customerId", app.UpdateHandler)
	r.Handler(http.MethodDelete, "/customers/:customerId", app.RemoveHandler)
	return r
}

// RetrieveCustomerID retrieves the customerID from an http request
func RetrieveCustomerID(r *http.Request) string {
	cID, err := retrieveCustomerID(r.Context())
	if err != nil {
		log.Fatal(err)
	}
	return cID
}

func retrieveCustomerID(ctx context.Context) (string, error) {
	params := httprouter.ParamsFromContext(ctx)
	cID := params.ByName("customerId")
	if cID == "" {
		return cID, errors.New("customerID not found in request context")
	}
	return cID, nil
}
