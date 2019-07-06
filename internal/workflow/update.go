package workflow

import (
	"github.com/albertowusuasare/customer-app/internal/msg"
	"github.com/albertowusuasare/customer-app/internal/storage"
	"github.com/albertowusuasare/customer-app/internal/updating"
)

type UpdateFunc func(request updating.Request) updating.UpdatedCustomer

func Update(updateDb storage.UpdateCustomerFunc, publishUpdate msg.CustomerUpdatedPublisherFunc) UpdateFunc {
	return func(r updating.Request) updating.UpdatedCustomer {
		updatedCustomer := updateDb(r)
		publishUpdate(updatedCustomer)
		return updatedCustomer
	}
}
