package google

import (
	"context"
	"log"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/albertowusuasare/customer-app/internal/adding"
	"github.com/albertowusuasare/customer-app/internal/retrieving"
	"github.com/albertowusuasare/customer-app/internal/storage"
	"github.com/albertowusuasare/customer-app/internal/uuid"
)

// CustomerDocument represents the firestore database entity for a customer
type CustomerDocument struct {
	CustomerID       string `firestore:"customerID"`
	FirstName        string `firestore:"firstName"`
	LastName         string `firestore:"lastName"`
	NationalID       string `firestore:"nationalID"`
	PhoneNumber      string `firestore:"phoneNumber"`
	AccountID        string `firestore:"accountID"`
	LastModifiedTime string `firestore:"lastModifiedTime"`
	CreatedTime      string `firestore:"createdTime"`
	Version          int    `firestore:"version"`
}

// CreateCustomerDoc inserts a customer document into google firestore
func CreateCustomerDoc(ctx context.Context, client *firestore.Client) storage.InsertCustomerFunc {
	return func(request adding.ValidatedRequest, genV4UUID uuid.GenV4Func) adding.Customer {
		v4UUID := genV4UUID()
		customerID := string(v4UUID)
		customerDoc := customerDocumentFromValidatedRequest(request, customerID)

		customers := client.Collection("Customer")
		customerDocRef := customers.Doc(customerID)
		_, err := customerDocRef.Create(ctx, customerDoc)

		if err != nil {
			log.Fatal(err)
		}

		return adding.Customer{
			CustomerID:  adding.CreateCustomerID(v4UUID),
			FirstName:   request.FirstName,
			LastName:    request.LastName,
			NationalID:  request.NationalID,
			PhoneNumber: request.PhoneNumber,
			AccountID:   request.AccountID,
		}
	}
}

func customerDocumentFromValidatedRequest(request adding.ValidatedRequest, customerID string) CustomerDocument {
	return CustomerDocument{
		CustomerID:       customerID,
		FirstName:        adding.RetrieveFirstName(request.FirstName),
		LastName:         adding.RetrieveLasttName(request.LastName),
		NationalID:       adding.RetrieveNationalID(request.NationalID),
		PhoneNumber:      adding.RetrievePhoneNumber(request.PhoneNumber),
		AccountID:        adding.RetrieveAccountID(request.AccountID),
		LastModifiedTime: time.Now().Format(time.RFC3339),
		CreatedTime:      time.Now().Format(time.RFC3339),
		Version:          0,
	}
}

// RetrieveCustomerDoc returns an in memory implementation of customer retrieval
func RetrieveCustomerDoc(ctx context.Context, client *firestore.Client) storage.RetrieveCustomerFunc {
	return func(customerID string) (*retrieving.Customer, error) {
		log.Printf("Retrieving customerId=%s from google firestore", customerID)

		customers := client.Collection("Customer")
		customerDocRef := customers.Doc(customerID)
		customerEntity, err := customerDocRef.Get(ctx)
		if err != nil {
			log.Printf("CustomerId=%s record does not exist in document store", customerID)
			return nil, retrieving.CustomerNonExistent{CustomerID: customerID}
		}

		var customerDoc CustomerDocument
		if err := customerEntity.DataTo(&customerDoc); err != nil {
			log.Fatal("Unable to convert customer entity to struct")
		}

		customer := customerFromCustomerDoc(customerDoc)
		return &customer, nil
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
