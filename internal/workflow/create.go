package workflow

import (
	"customer-app/internal/adding"
	"customer-app/internal/msg"
	"customer-app/internal/storage"
	"customer-app/internal/uuid"
)

type CreateFunc func(r adding.Request) adding.PersistedCustomer

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
