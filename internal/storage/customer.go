package storage

import (
	"github.com/albertowusuasare/customer-app/internal/adding"
	"github.com/albertowusuasare/customer-app/internal/retrieving"
	"github.com/albertowusuasare/customer-app/internal/updating"
	"github.com/albertowusuasare/customer-app/internal/uuid"
)

// InsertCustomerFunc adds a new customer to a permaent data store.
// A unique id for the customer will be generated using genUUIDStr
type InsertCustomerFunc func(unPersistedCustomer adding.UnPersistedCustomer, genUUIDStr uuid.GenFunc) adding.PersistedCustomer

// RetrieveCustomerFunc retrieve a customer in the data store corresponding to custoemrId
type RetrieveCustomerFunc func(customerId string) retrieving.Customer

// RetrieveCustomersFunc retrieves multiple customers based on the query criteria defined in 'request'.
type RetrieveCustomersFunc func(request retrieving.MultiRequest) []retrieving.Customer

// UpdateCustomerFunc updates a customer in the datastore based on the incoming request
type UpdateCustomerFunc func(request updating.Request) updating.UpdatedCustomer

// RemoveCustomerFunc removes the customer corresponding to customerId from the datastore
type RemoveCustomerFunc func(customerId string)
