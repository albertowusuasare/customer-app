package workflow

import (
	"customer-app/internal/msg"
	"customer-app/internal/storage"
)

type RemoveFunc func(customerId string)

func Remove(removeCustomer storage.RemoveCustomerFunc, publishRemove msg.CustomerRemovedPublisherFunc) RemoveFunc {
	return func(customerId string) {
		removeCustomer(customerId)
		publishRemove(customerId)
	}
}
