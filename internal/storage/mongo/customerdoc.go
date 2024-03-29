package mongo

import (
	"github.com/albertowusuasare/customer-app/internal/adding"
	"github.com/albertowusuasare/customer-app/internal/retrieving"
	"github.com/albertowusuasare/customer-app/internal/storage"
	"github.com/albertowusuasare/customer-app/internal/updating"
)

//InsertCustomer returns a mongo implementation for customer inserts
func InsertCustomer() storage.InsertCustomerFunc {
	return func(c *adding.Customer) error {
		panic("Mongo insert not implemented yet")
	}
}

// RetrieveCustomer returns a mongo implementation of customer retrieval
func RetrieveCustomer() storage.RetrieveCustomerFunc {
	return func(customerID string) (*retrieving.Customer, error) {
		panic("Mongo retrieve not implemented yet")
	}
}

// RetrieveCustomers returns a mongo implementation of customers retrieval
func RetrieveCustomers() storage.RetrieveCustomersFunc {
	return func(request retrieving.MultiRequest) []retrieving.Customer {
		panic("Mongo retrieve all not implemented yet")
	}
}

// UpdateCustomer returns a mongo implementation for customer updates
func UpdateCustomer() storage.UpdateCustomerFunc {
	return func(request updating.Request) updating.UpdatedCustomer {
		panic("Mongo update not implemented yet")
	}
}

//RemoveCustomer returns a mongo  implementation of customer removal
func RemoveCustomer() storage.RemoveCustomerFunc {
	return func(customerID string) {
		panic("Mongo remove not implemented yet")
	}
}
