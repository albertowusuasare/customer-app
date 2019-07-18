package storage

import (
	"github.com/albertowusuasare/customer-app/internal/adding"
	"github.com/albertowusuasare/customer-app/internal/retrieving"
	"github.com/albertowusuasare/customer-app/internal/updating"
	"github.com/albertowusuasare/customer-app/internal/uuid"
)

// InsertCustomerFunc adds a new customer to a permaent data store.
// A unique id for the customer will be generated using genUUIDStr
type InsertCustomerFunc func(request adding.ValidatedRequest, genUUIDStr uuid.GenFunc) adding.PersistedCustomer

// RetrieveCustomerFunc retrieve a customer in the data store corresponding to customerId
// when error is nil. An error is returned if a customer record cannot be found for customerId.
// In that case retrieving.CustomerNonExistent is returned
type RetrieveCustomerFunc func(customerId string) (*retrieving.Customer, error)

// RetrieveCustomersFunc retrieves multiple customers based on the query criteria defined in 'request'.
type RetrieveCustomersFunc func(request retrieving.MultiRequest) []retrieving.Customer

// UpdateCustomerFunc updates a customer in the datastore based on the incoming request
type UpdateCustomerFunc func(request updating.Request) updating.UpdatedCustomer

// RemoveCustomerFunc removes the customer corresponding to customerId from the datastore
type RemoveCustomerFunc func(customerId string)
