package handler

import (
	"net/http"

	"github.com/albertowusuasare/customer-app/internal/workflow"
)

// RemoveHandler represents the http handler for a customer remove http call
type RemoveHandler struct {
	Workflow workflow.RemoveFunc
}

// Handle allows the RemoveHandler to act as an http call handler
func (h RemoveHandler) Handle(w http.ResponseWriter, r *http.Request) {
	customerId := RetrieveCustomerId(r)
	h.Workflow(customerId)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(204)
}
