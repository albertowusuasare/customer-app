package mongo

import (
	"customer-app/internal/adding"
	"customer-app/internal/retrieving"
	"customer-app/internal/storage"
	"customer-app/internal/updating"
	"customer-app/internal/uuid"
)

func InsertCustomer() storage.InsertCustomerFunc {
	return func(unPersistedCustomer adding.UnPersistedCustomer, genUUIDStr uuid.GenFunc) adding.PersistedCustomer {
		panic("Mongo insert not implemented yet")
	}
}

func RetrieveCustomer() storage.RetrieveCustomerFunc {
	return func(customerId string) retrieving.Customer {
		panic("Mongo retrieve not implemented yet")
	}
}

func RetrieveCustomers() storage.RetrieveCustomersFunc {
	return func(request retrieving.MultiRequest) []retrieving.Customer {
		panic("Mongo retrieve all not implemented yet")
	}
}

func UpdateCustomer() storage.UpdateCustomerFunc {
	return func(request updating.Request) updating.UpdatedCustomer {
		panic("Mongo update not implemented yet")
	}
}

func RemoveCustomer() storage.RemoveCustomerFunc {
	return func(customerId string) {
		panic("Mongo remove not implemented yet")
	}
}
