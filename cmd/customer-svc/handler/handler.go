package handler

import (
	"customer-app/internal/app"
	"net/http"
)

func Handle(app app.Customer) http.Handler {
	createHandler := CreateHandler{Workflow: app.CreateWf}
	retrieveOneHandler := RetrieveOneHandler{Workflow: app.RetrieveSingleWf}
	retrieveMultiHandler := RetrieveMultiHandler{Workflow: app.RetrieveMultiWf}
	updateHandler := UpdateHandler{Workflow: app.UpdateWf}
	removeHandler := RemoveHandler{Workflow: app.RemoveWf}
	return Router(createHandler, retrieveOneHandler, retrieveMultiHandler, updateHandler, removeHandler)
}
