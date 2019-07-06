package workflow

import (
	"customer-app/internal/retrieving"
	"customer-app/internal/storage"
	"reflect"
	"testing"
)

func TestRetrieveOne(t *testing.T) {
	customerId := "fb57de5c-8679-476c-a546-6304a9fe5bb3"
	expectedCustomer := retrieveFromDb(customerId)
	retrieveOne := RetrieveOne(retrieveFromDb)
	actualCustomer := retrieveOne(customerId)

	if expectedCustomer != actualCustomer {
		t.Errorf("with customerId=%s, expectedCustomer=%+v actualCustomer=%+v",
			customerId, expectedCustomer, actualCustomer)
	}
}

func retrieveFromDb(customerId string) retrieving.Customer {
	return mockCustomer(customerId)
}

func mockCustomer(customerId string) retrieving.Customer {
	return retrieving.Customer{
		CustomerId:       customerId,
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
