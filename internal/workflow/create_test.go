package workflow

import (
	"customer-app/internal/adding"
	"customer-app/internal/msg"
	"customer-app/internal/storage"
	"customer-app/internal/uuid"
	"fmt"
	"testing"
)

func TestCreate(t *testing.T) {
	request := mockAddingRequest()
	expectedId := "724377e7-1567-4756-b18b-1a15ed35d8f4"
	expectedPersistedCustomer := newPersistedCustomer(expectedId, request)

	requestValidator := successRequestValidator()
	genUUIDStr := func() string { return expectedId }
	insertCustomer := successInsertCustomer()
	publishCustomerAdded := successCustomerAddedPublisher(expectedPersistedCustomer, t)

	createFunc := Create(requestValidator, genUUIDStr, insertCustomer, publishCustomerAdded)
	actualPersistedCustomer := createFunc(request)

	if expectedPersistedCustomer != actualPersistedCustomer {
		t.Errorf("expectedPersistedCustomer=%+v is not equal to actualPersistedCustomer=%+v", expectedPersistedCustomer, actualPersistedCustomer)
	}

}

func mockAddingRequest() adding.Request {
	return adding.Request{
		FirstName:   "John",
		LastName:    "Doe",
		NationalId:  "987654321",
		PhoneNumber: "020765432",
		AccountId:   "b253fd2a-6bb9-49db-9fb5-f6388e1661a7",
	}
}

func newPersistedCustomer(id string, r adding.Request) adding.PersistedCustomer {
	return adding.PersistedCustomer{
		CustomerId:  id,
		FirstName:   r.FirstName,
		LastName:    r.LastName,
		NationalId:  r.NationalId,
		PhoneNumber: r.PhoneNumber,
		AccountId:   r.AccountId,
	}
}

func successRequestValidator() adding.RequestValidatorFunc {
	return func(r adding.Request) (adding.UnPersistedCustomer, error) {
		return adding.UnPersistedCustomer{
			FirstName:   r.FirstName,
			LastName:    r.LastName,
			NationalId:  r.NationalId,
			PhoneNumber: r.PhoneNumber,
			AccountId:   r.AccountId,
		}, nil
	}
}

func successInsertCustomer() storage.InsertCustomerFunc {
	return func(unPersistedCustomer adding.UnPersistedCustomer, genUUIDStr uuid.GenFunc) adding.PersistedCustomer {
		return adding.PersistedCustomer{
			CustomerId:  genUUIDStr(),
			FirstName:   unPersistedCustomer.FirstName,
			LastName:    unPersistedCustomer.LastName,
			NationalId:  unPersistedCustomer.NationalId,
			PhoneNumber: unPersistedCustomer.PhoneNumber,
			AccountId:   unPersistedCustomer.AccountId,
		}
	}
}

func successCustomerAddedPublisher(expectedPersistedCustomer adding.PersistedCustomer,
	t *testing.T) msg.CustomerAddedPublisherFunc {
	return func(customer adding.PersistedCustomer) msg.Response {
		if expectedPersistedCustomer != customer {
			t.Errorf(fmt.Sprintf("Invalid customerAddedPublish arg. Expected=%+v Actual=%+v", expectedPersistedCustomer, customer))
		}
		return msg.Response{
			MessageId:    "19a1eed3-a650-412d-aeb7-20fabe0b37bc",
			Acknowledged: true,
		}
	}

}
