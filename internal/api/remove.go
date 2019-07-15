package api

import (
	"net/http"

	"github.com/albertowusuasare/customer-app/internal/workflow"
)

// HandleRemove returns an http handler for a customer remove API call
func HandleRemove(wf workflow.RemoveFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		customerId := RetrieveCustomerId(r)
		wf(customerId)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(204)
	}
}
