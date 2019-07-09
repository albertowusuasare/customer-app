package workflow

import (
	"reflect"
	"testing"

	"github.com/albertowusuasare/customer-app/internal/retrieving"
	"github.com/albertowusuasare/customer-app/internal/storage"
)

func TestRetrieveOne(t *testing.T) {
	customerID := "fb57de5c-8679-476c-a546-6304a9fe5bb3"
	expectedCustomer, _ := retrieveFromDb(customerID)
	retrieveOne := RetrieveOne(retrieveFromDb)
	actualCustomer, _ := retrieveOne(customerID)

	if *expectedCustomer != *actualCustomer {
		t.Errorf("with customerId=%s, expectedCustomer=%+v actualCustomer=%+v",
			customerID, *expectedCustomer, *actualCustomer)
	}
}

func retrieveFromDb(customerID string) (*retrieving.Customer, error) {
	customer := mockCustomer(customerID)
	return &customer, nil
}

func mockCustomer(customerID string) retrieving.Customer {
	return retrieving.Customer{
		CustomerId:       customerID,
		FirstName:        "John",
		LastName:         "Doe",
		NationalId:       "9876543",
		PhoneNumber:      "987654321",
		AccountId:        "1bd23762-b26e-4bdb-8203-b4f1eecc003d",
		LastModifiedTime: "2019-07-05T01:39:36+01:00",
		CreatedTime:      "2019-07-05T01:39:36+01:00",
		Version:          0,
	}
}

func TestRetrieveOneError(t *testing.T) {
	customerID := "45dca513-c40f-4f9c-b7b4-1b2e385b343d"
	expectedError := retrieving.CustomerNonExistent{CustomerID: customerID}
	retrieveFromDb := func(cID string) (*retrieving.Customer, error) {
		return nil, expectedError
	}
	retrieveOne := RetrieveOne(retrieveFromDb)
	_, err := retrieveOne(customerID)

	if expectedError != err {
		t.Errorf("Expecting error %s but got %s", expectedError, err)
	}
}

func TestRetrieveMulti(t *testing.T) {
	request := retrieving.MultiRequest{}
	retrieveDbCustomers := retrieveDbCustomers(request, t)

	retrieveMulti := RetrieveMulti(retrieveDbCustomers)
	expected := retrieveDbCustomers(request)
	actual := retrieveMulti(request)

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("epectedCustomers=%+v is not equal to actualCustomers=%+v", expected, actual)
	}
}

func retrieveDbCustomers(
	expectedRequest retrieving.MultiRequest,
	t *testing.T) storage.RetrieveCustomersFunc {
	return func(request retrieving.MultiRequest) []retrieving.Customer {
		if !reflect.DeepEqual(expectedRequest, request) {
			t.Errorf("expectedRequest=%+v is not equal to actualRequest=%+v", expectedRequest, request)
		}
		customers := []retrieving.Customer{}
		customer := mockCustomer("0b281b75-b0db-4e66-a258-41f97458ec04")
		return append(customers, customer)
	}

}
