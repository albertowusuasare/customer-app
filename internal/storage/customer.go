package storage

import (
	"github.com/albertowusuasare/customer-app/internal/adding"
	"github.com/albertowusuasare/customer-app/internal/retrieving"
	"github.com/albertowusuasare/customer-app/internal/updating"
)

// InsertCustomerFunc adds a new customer to a permaent data store.
// An error is returned if an unexpected failure happens during the insert
type InsertCustomerFunc func(c *adding.Customer) error

// RetrieveCustomerFunc retrieve a customer in the data store corresponding to customerID
// when error is nil. An error is returned if a customer record cannot be found for customerID.
// In that case retrieving.CustomerNonExistent is returned
type RetrieveCustomerFunc func(customerID string) (*retrieving.Customer, error)

// RetrieveCustomersFunc retrieves multiple customers based on the query criteria defined in 'request'.
type RetrieveCustomersFunc func(request retrieving.MultiRequest) []retrieving.Customer

// UpdateCustomerFunc updates a customer in the datastore based on the incoming request
type UpdateCustomerFunc func(request updating.Request) updating.UpdatedCustomer

// RemoveCustomerFunc removes the customer corresponding to customerID from the datastore
type RemoveCustomerFunc func(customerID string)
