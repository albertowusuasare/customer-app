package workflow

import (
	"github.com/albertowusuasare/customer-app/internal/adding"
	"github.com/albertowusuasare/customer-app/internal/msg"
	"github.com/albertowusuasare/customer-app/internal/storage"
	"github.com/albertowusuasare/customer-app/internal/uuid"
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
