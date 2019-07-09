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

// CustomerDocument represents the database entity for a customer
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

//InsertCustomer returns an imemory implementation for customer inserts
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

// RetrieveCustomer returns an in memory implementation of customer retrieval
func RetrieveCustomer() storage.RetrieveCustomerFunc {
	return func(customerID string) (*retrieving.Customer, error) {
		log.Printf("Retrieving customerId=%s from the in memory database", customerID)
		customerDoc, present := customerCollection[customerID]
		if present {
			customer := customerFromCustomerDoc(customerDoc)
			return &customer, nil
		}
		log.Printf("CustomerId=%s record does not exist in document store", customerID)
		return nil, retrieving.CustomerNonExistent{CustomerID: customerID}
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

// RetrieveCustomers returns an in memory implementation of customers retrieval
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

// UpdateCustomer returns an in memory implementation for customer updates
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

//RemoveCustomer returns an in memory implementation of customer removal
func RemoveCustomer() storage.RemoveCustomerFunc {
	return func(customerId string) {
		delete(customerCollection, customerId)
	}
}
