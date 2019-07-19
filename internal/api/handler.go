package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/albertowusuasare/customer-app/internal/app"
	"github.com/google/uuid"
)

// HandlerFunc returns an entry point http handler for the entire application
type HandlerFunc func(app app.Customer) http.Handler

// Handler is the default http handler for the application
func Handler(app app.Customer) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler, err := resolveHandler(app, w, r)
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
		handler.ServeHTTP(w, r)
	})
}

func resolveHandler(app app.Customer, w http.ResponseWriter, r *http.Request) (http.Handler, error) {
	method := r.Method
	if method == "POST" {
		return HandleCreate(app.CreateWf), nil
	}

	if method == "GET" {
		uri := r.RequestURI
		if hasCustomerID(uri) {
			return HandleRetrieveOne(app.RetrieveSingleWf), nil
		}
		return HandleRetrieveMulti(app.RetrieveMultiWf), nil

	}

	if method == "PUT" {
		return HandleUpdate(app.UpdateWf), nil
	}

	if method == "DELETE" {
		return HandleRemove(app.RemoveWf), nil
	}

	return nil, fmt.Errorf("httpMethod=%s not supported", r.Method)

}

func hasCustomerID(requestURI string) bool {
	value := strings.Split(requestURI, "/")[2]
	_, err := uuid.Parse(value)
	return err == nil
}

// RetrieveCustomerID retrieves the customerID from an http request
func RetrieveCustomerID(r *http.Request) string {
	return strings.Split(r.RequestURI, "/")[2]
}
