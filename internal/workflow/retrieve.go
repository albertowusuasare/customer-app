package workflow

import (
	"github.com/albertowusuasare/customer-app/internal/retrieving"
	"github.com/albertowusuasare/customer-app/internal/storage"
)

// RetrieveSingleFunc retrieves an existing customer with customerID from the datastore
// An error is returned when there is no corresponding customer record for customerID
// Specifically, retrieving.CustomerNonExistent will be returned in that case.
type RetrieveSingleFunc func(customerID string) (*retrieving.Customer, error)

// RetrieveMultiFunc retrieves multiple customers that match the query describe in r.
type RetrieveMultiFunc func(r retrieving.MultiRequest) []retrieving.Customer

// RetrieveOne is the default implementation of the customer retrieve workflow
func RetrieveOne(retrieveFromDb storage.RetrieveCustomerFunc) RetrieveSingleFunc {
	return func(customerID string) (*retrieving.Customer, error) {
		return retrieveFromDb(customerID)
	}
}

// RetrieveMulti is the default implementation of the customers retrieve workflow
func RetrieveMulti(retrieveCustomers storage.RetrieveCustomersFunc) RetrieveMultiFunc {
	return func(r retrieving.MultiRequest) []retrieving.Customer {
		return retrieveCustomers(r)
	}
}
