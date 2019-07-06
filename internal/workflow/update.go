package workflow

import (
	"customer-app/internal/msg"
	"customer-app/internal/storage"
	"customer-app/internal/updating"
)

type UpdateFunc func(request updating.Request) updating.UpdatedCustomer

func Update(updateDb storage.UpdateCustomerFunc, publishUpdate msg.CustomerUpdatedPublisherFunc) UpdateFunc {
	return func(r updating.Request) updating.UpdatedCustomer {
		updatedCustomer := updateDb(r)
		publishUpdate(updatedCustomer)
		return updatedCustomer
	}
}
