package workflow

import (
	"github.com/albertowusuasare/customer-app/internal/msg"
	"github.com/albertowusuasare/customer-app/internal/storage"
)

type RemoveFunc func(customerId string)

func Remove(removeCustomer storage.RemoveCustomerFunc, publishRemove msg.CustomerRemovedPublisherFunc) RemoveFunc {
	return func(customerId string) {
		removeCustomer(customerId)
		publishRemove(customerId)
	}
}
