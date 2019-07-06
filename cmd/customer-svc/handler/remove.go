package handler

import (
	"net/http"

	"github.com/albertowusuasare/customer-app/internal/workflow"
)

type RemoveHandler struct {
	Workflow workflow.RemoveFunc
}

func (h RemoveHandler) Handle(w http.ResponseWriter, r *http.Request) {
	customerId := RetrieveCustomerId(r)
	h.Workflow(customerId)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(204)
}
