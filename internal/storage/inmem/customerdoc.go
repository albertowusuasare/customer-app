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
	CustomerID       string
	FirstName        string
	LastName         string
	NationalID       string
	PhoneNumber      string
	AccountID        string
	LastModifiedTime string
	CreatedTime      string
	Version          int
}

var customerCollection = map[string]CustomerDocument{}

//InsertCustomer returns an imemory implementation for customer inserts
func InsertCustomer() storage.InsertCustomerFunc {
	return func(request adding.ValidatedRequest, genUUIDStr uuid.GenFunc) adding.Customer {
		customerDoc := customerDocumentFromValidatedRequest(request, genUUIDStr)
		fmt.Printf("Adding customerDoc=%+v to in memory database\n", customerDoc)
		customerCollection[customerDoc.CustomerID] = customerDoc
		return adding.Customer{
			CustomerID:  customerDoc.CustomerID,
			FirstName:   customerDoc.FirstName,
			LastName:    customerDoc.LastName,
			NationalID:  customerDoc.NationalID,
			PhoneNumber: customerDoc.PhoneNumber,
			AccountID:   customerDoc.AccountID,
		}
	}
}

func customerDocumentFromValidatedRequest(request adding.ValidatedRequest, genUUIDStr uuid.GenFunc) CustomerDocument {
	return CustomerDocument{
		CustomerID:       genUUIDStr(),
		FirstName:        adding.RetrieveFirstName(request.FirstName),
		LastName:         request.LastName,
		NationalID:       request.NationalID,
		PhoneNumber:      request.PhoneNumber,
		AccountID:        request.AccountID,
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
		CustomerID:       customerDoc.CustomerID,
		FirstName:        customerDoc.FirstName,
		LastName:         customerDoc.LastName,
		NationalID:       customerDoc.NationalID,
		PhoneNumber:      customerDoc.PhoneNumber,
		AccountID:        customerDoc.AccountID,
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
		priorDocument := customerCollection[request.CustomerID]

		customerDoc := CustomerDocument{
			CustomerID:       priorDocument.CustomerID,
			FirstName:        request.FirstName,
			LastName:         request.LastName,
			NationalID:       request.NationalID,
			PhoneNumber:      request.PhoneNumber,
			AccountID:        priorDocument.AccountID,
			LastModifiedTime: time.Now().Format(time.RFC3339),
			CreatedTime:      priorDocument.CreatedTime,
			Version:          priorDocument.Version + 1,
		}

		customerCollection[customerDoc.CustomerID] = customerDoc

		return updating.UpdatedCustomer{
			CustomerID:       customerDoc.CustomerID,
			FirstName:        customerDoc.FirstName,
			LastName:         customerDoc.LastName,
			NationalID:       customerDoc.NationalID,
			PhoneNumber:      customerDoc.PhoneNumber,
			AccountID:        customerDoc.AccountID,
			LastModifiedTime: customerDoc.LastModifiedTime,
			CreatedTime:      customerDoc.CreatedTime,
			Version:          customerDoc.Version,
		}

	}
}

//RemoveCustomer returns an in memory implementation of customer removal
func RemoveCustomer() storage.RemoveCustomerFunc {
	return func(customerID string) {
		delete(customerCollection, customerID)
	}
}
