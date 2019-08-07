package app

import (
	"net/http"

	"github.com/albertowusuasare/customer-app/internal/workflow"
)

// HandleRemove returns an http handler for a customer remove API call
func HandleRemove(wf workflow.RemoveFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		customerID := RetrieveCustomerID(r)
		wf(customerID)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(204)
	}
}
