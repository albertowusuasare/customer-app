package app

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"
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
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler, err := resolveHandler(app, w, r)
		if err != nil {
			http.Error(w, err.Error(), 500)
		}
		handler.ServeHTTP(w, r)
	})
}

func resolveHandler(app StandAlone, w http.ResponseWriter, r *http.Request) (http.Handler, error) {
	method := r.Method
	if method == "POST" {
		return app.CreateHandler, nil
	}

	if method == "GET" {
		uri := r.RequestURI
		if hasCustomerID(uri) {
			return app.RetrieveOneHandler, nil
		}
		return app.RetrieveMultiHandler, nil

	}

	if method == "PUT" {
		return app.UpdateHandler, nil
	}

	if method == "DELETE" {
		return app.RemoveHandler, nil
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
