package workflow

import (
	"github.com/albertowusuasare/customer-app/internal/msg"
	"github.com/albertowusuasare/customer-app/internal/storage"
)

// RemoveFunc removes a customer from the datastore
type RemoveFunc func(customerId string)

// Remove is the default implementation of the customer removal workflow
func Remove(removeCustomer storage.RemoveCustomerFunc, publishRemove msg.CustomerRemovedPublisherFunc) RemoveFunc {
	return func(customerId string) {
		removeCustomer(customerId)
		publishRemove(customerId)
	}
}
