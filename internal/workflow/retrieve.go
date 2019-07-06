package workflow

import (
	"github.com/albertowusuasare/customer-app/internal/retrieving"
	"github.com/albertowusuasare/customer-app/internal/storage"
)

type RetrieveSingleFunc func(customerId string) retrieving.Customer

type RetrieveMultiFunc func(r retrieving.MultiRequest) []retrieving.Customer

func RetrieveOne(retrieveFromDb storage.RetrieveCustomerFunc) RetrieveSingleFunc {
	return func(customerId string) retrieving.Customer {
		return retrieveFromDb(customerId)
	}
}

func RetrieveMulti(retrieveCustomers storage.RetrieveCustomersFunc) RetrieveMultiFunc {
	return func(r retrieving.MultiRequest) []retrieving.Customer {
		return retrieveCustomers(r)
	}
}
