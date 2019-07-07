package pkg

import (
	"github.com/albertowusuasare/customer-app/internal/adding"
	"github.com/albertowusuasare/customer-app/internal/app"
	queue "github.com/albertowusuasare/customer-app/internal/msg/inmem"
	"github.com/albertowusuasare/customer-app/internal/storage/inmem"
	"github.com/albertowusuasare/customer-app/internal/uuid"
	"github.com/albertowusuasare/customer-app/internal/workflow"
)

// CustomerAppFunc creates a new instance of the customer application
type CustomerAppFunc func() app.Customer

// InmemApp creates a customer app based on in memory data store
func InmemApp() app.Customer {
	createWf := workflow.Create(adding.RequestToUnPersistedCustomer, uuid.Gen(), inmem.InsertCustomer(), queue.CustomerAddedPublisher())
	retrieveSingleWf := workflow.RetrieveOne(inmem.RetrieveCustomer())
	retrieveMultiWf := workflow.RetrieveMulti(inmem.RetrieveCustomers())
	updateWf := workflow.Update(inmem.UpdateCustomer(), queue.CustomerUpdatedPublisher())
	removeWf := workflow.Remove(inmem.RemoveCustomer(), queue.CustomerRemovedPublisher())

	return app.Customer{
		CreateWf:         createWf,
		RetrieveSingleWf: retrieveSingleWf,
		RetrieveMultiWf:  retrieveMultiWf,
		UpdateWf:         updateWf,
		RemoveWf:         removeWf,
	}
}
