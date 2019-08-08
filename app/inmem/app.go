package inmem

import (
	"github.com/albertowusuasare/customer-app/app"
	"github.com/albertowusuasare/customer-app/internal/adding"
	queue "github.com/albertowusuasare/customer-app/internal/msg/inmem"
	"github.com/albertowusuasare/customer-app/internal/storage/inmem"
	"github.com/albertowusuasare/customer-app/internal/uuid"
	"github.com/albertowusuasare/customer-app/internal/workflow"
)

// App creates a stand alone app based on in memory data store
func App() app.StandAlone {
	createWf := workflow.Create(adding.ValidateRequest, uuid.GenV4, inmem.InsertCustomer(), queue.CustomerAddedPublisher())
	createHandler := app.HandleCreate(createWf)

	retrieveSingleWf := workflow.RetrieveOne(inmem.RetrieveCustomer())
	retrieveOneHandler := app.HandleRetrieveOne(retrieveSingleWf)

	retrieveMultiWf := workflow.RetrieveMulti(inmem.RetrieveCustomers())
	retrieveMultiHandler := app.HandleRetrieveMulti(retrieveMultiWf)

	updateWf := workflow.Update(inmem.UpdateCustomer(), queue.CustomerUpdatedPublisher())
	updateHandler := app.HandleUpdate(updateWf)

	removeWf := workflow.Remove(inmem.RemoveCustomer(), queue.CustomerRemovedPublisher())
	removeHandler := app.HandleRemove(removeWf)

	return app.StandAlone{
		CreateHandler:        createHandler,
		RetrieveOneHandler:   retrieveOneHandler,
		RetrieveMultiHandler: retrieveMultiHandler,
		UpdateHandler:        updateHandler,
		RemoveHandler:        removeHandler,
	}
}
