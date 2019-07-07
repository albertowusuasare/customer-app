package workflow

import (
	"github.com/albertowusuasare/customer-app/internal/msg"
	"github.com/albertowusuasare/customer-app/internal/storage"
	"github.com/albertowusuasare/customer-app/internal/updating"
)

// UpdateFunc updates a customer in the datastore based on request
type UpdateFunc func(request updating.Request) updating.UpdatedCustomer

// Update is the default implementation of the customer update workflow
func Update(updateDb storage.UpdateCustomerFunc, publishUpdate msg.CustomerUpdatedPublisherFunc) UpdateFunc {
	return func(r updating.Request) updating.UpdatedCustomer {
		updatedCustomer := updateDb(r)
		publishUpdate(updatedCustomer)
		return updatedCustomer
	}
}
