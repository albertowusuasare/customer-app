package inmem

import (
	"fmt"
	"log"
	"time"

	"github.com/albertowusuasare/customer-app/internal/adding"
	"github.com/albertowusuasare/customer-app/internal/retrieving"
	"github.com/albertowusuasare/customer-app/internal/storage"
	"github.com/albertowusuasare/customer-app/internal/updating"
	"github.com/albertowusuasare/customer-app/internal/uuid"
)

//Collection
type CustomerDocument struct {
	CustomerId       string
	FirstName        string
	LastName         string
	NationalId       string
	PhoneNumber      string
	AccountId        string
	LastModifiedTime string
	CreatedTime      string
	Version          int
}

var customerCollection = map[string]CustomerDocument{}

//Create
func InsertCustomer() storage.InsertCustomerFunc {
	return func(customer adding.UnPersistedCustomer, genUUIDStr uuid.GenFunc) adding.PersistedCustomer {
		customerDoc := customerDocumentFromUnPersistedCustomer(customer, genUUIDStr)
		fmt.Printf("Adding customerDoc=%+v to in memory database\n", customerDoc)
		customerCollection[customerDoc.CustomerId] = customerDoc
		return adding.PersistedCustomer{
			CustomerId:  customerDoc.CustomerId,
			FirstName:   customerDoc.FirstName,
			LastName:    customerDoc.LastName,
			NationalId:  customerDoc.NationalId,
			PhoneNumber: customerDoc.PhoneNumber,
			AccountId:   customerDoc.AccountId,
		}
	}
}

func customerDocumentFromUnPersistedCustomer(customer adding.UnPersistedCustomer, genUUIDStr uuid.GenFunc) CustomerDocument {
	return CustomerDocument{
		CustomerId:       genUUIDStr(),
		FirstName:        customer.FirstName,
		LastName:         customer.LastName,
		NationalId:       customer.NationalId,
		PhoneNumber:      customer.PhoneNumber,
		AccountId:        customer.AccountId,
		LastModifiedTime: time.Now().Format(time.RFC3339),
		CreatedTime:      time.Now().Format(time.RFC3339),
		Version:          0,
	}
}

//Retrieve Customer
func RetrieveCustomer() storage.RetrieveCustomerFunc {
	return func(customerId string) retrieving.Customer {
		log.Printf("Retrieving customerId=%s from the in memory database", customerId)
		customerDoc := customerCollection[customerId]
		return customerFromCustomerDoc(customerDoc)
	}
}

func customerFromCustomerDoc(customerDoc CustomerDocument) retrieving.Customer {
	return retrieving.Customer{
		CustomerId:       customerDoc.CustomerId,
		FirstName:        customerDoc.FirstName,
		LastName:         customerDoc.LastName,
		NationalId:       customerDoc.NationalId,
		PhoneNumber:      customerDoc.PhoneNumber,
		AccountId:        customerDoc.AccountId,
		LastModifiedTime: customerDoc.LastModifiedTime,
		CreatedTime:      customerDoc.CreatedTime,
		Version:          customerDoc.Version,
	}
}

func RetrieveCustomers() storage.RetrieveCustomersFunc {
	return func(request retrieving.MultiRequest) []retrieving.Customer {
		customers := []retrieving.Customer{}
		for _, v := range customerCollection {
			customerDoc := customerFromCustomerDoc(v)
			customers = append(customers, customerDoc)
		}
		return customers
	}
}

//Update Customer
func UpdateCustomer() storage.UpdateCustomerFunc {
	return func(request updating.Request) updating.UpdatedCustomer {
		piorDocument := customerCollection[request.CustomerId]

		customerDoc := CustomerDocument{
			CustomerId:       piorDocument.CustomerId,
			FirstName:        request.FirstName,
			LastName:         request.LastName,
			NationalId:       request.NationalId,
			PhoneNumber:      request.PhoneNumber,
			AccountId:        piorDocument.AccountId,
			LastModifiedTime: time.Now().Format(time.RFC3339),
			CreatedTime:      piorDocument.CreatedTime,
			Version:          piorDocument.Version + 1,
		}

		customerCollection[customerDoc.CustomerId] = customerDoc

		return updating.UpdatedCustomer{
			CustomerId:       customerDoc.CustomerId,
			FirstName:        customerDoc.FirstName,
			LastName:         customerDoc.LastName,
			NationalId:       customerDoc.NationalId,
			PhoneNumber:      customerDoc.PhoneNumber,
			AccountId:        customerDoc.AccountId,
			LastModifiedTime: customerDoc.LastModifiedTime,
			CreatedTime:      customerDoc.CreatedTime,
			Version:          customerDoc.Version,
		}

	}
}

//Remove Customer
func RemoveCustomer() storage.RemoveCustomerFunc {
	return func(customerId string) {
		delete(customerCollection, customerId)
	}
}
