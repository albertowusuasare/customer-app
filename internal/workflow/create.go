package workflow

import (
	"github.com/albertowusuasare/customer-app/internal/adding"
	"github.com/albertowusuasare/customer-app/internal/msg"
	"github.com/albertowusuasare/customer-app/internal/storage"
	"github.com/albertowusuasare/customer-app/internal/uuid"
)

// CreateFunc creates a persisted customer from a request to add a new customer
type CreateFunc func(r adding.Request) adding.PersistedCustomer

// Create is the default implementation of the customer create workflow
func Create(
	validateRequest adding.RequestValidatorFunc,
	genUUIDStr uuid.GenFunc,
	insertCustomer storage.InsertCustomerFunc,
	publishCustomerAdded msg.CustomerAddedPublisherFunc) CreateFunc {
	return func(request adding.Request) adding.PersistedCustomer {
		unPersistedCustomer, _ := validateRequest(request)
		persistedCustomer := insertCustomer(unPersistedCustomer, genUUIDStr)
		publishCustomerAdded(persistedCustomer)
		return persistedCustomer
	}
}
