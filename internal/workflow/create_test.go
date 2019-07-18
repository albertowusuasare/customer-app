package workflow

import (
	"fmt"
	"testing"

	"github.com/albertowusuasare/customer-app/internal/adding"
	"github.com/albertowusuasare/customer-app/internal/msg"
	"github.com/albertowusuasare/customer-app/internal/storage"
	"github.com/albertowusuasare/customer-app/internal/uuid"
)

func TestCreate(t *testing.T) {
	request := mockAddingRequest()
	expectedID := "724377e7-1567-4756-b18b-1a15ed35d8f4"
	expectedPersistedCustomer := newPersistedCustomer(expectedID, request)

	requestValidator := successRequestValidator()
	genUUIDStr := func() string { return expectedID }
	insertCustomer := successInsertCustomer()
	publishCustomerAdded := successCustomerAddedPublisher(expectedPersistedCustomer, t)

	createFunc := Create(requestValidator, genUUIDStr, insertCustomer, publishCustomerAdded)
	actualPersistedCustomer, _ := createFunc(request)

	if expectedPersistedCustomer != actualPersistedCustomer {
		t.Errorf("expectedPersistedCustomer=%+v is not equal to actualPersistedCustomer=%+v", expectedPersistedCustomer, actualPersistedCustomer)
	}

}

func mockAddingRequest() adding.UnvalidatedRequest {
	return adding.UnvalidatedRequest{
		FirstName:   "John",
		LastName:    "Doe",
		NationalId:  "987654321",
		PhoneNumber: "020765432",
		AccountId:   "b253fd2a-6bb9-49db-9fb5-f6388e1661a7",
	}
}

func newPersistedCustomer(id string, r adding.UnvalidatedRequest) adding.PersistedCustomer {
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
	return func(r adding.UnvalidatedRequest) (adding.ValidatedRequest, error) {
		firstName, _ := adding.CreateFirstName(r.FirstName)
		return adding.ValidatedRequest{
			FirstName:   firstName,
			LastName:    r.LastName,
			NationalId:  r.NationalId,
			PhoneNumber: r.PhoneNumber,
			AccountId:   r.AccountId,
		}, nil
	}
}

func successInsertCustomer() storage.InsertCustomerFunc {
	return func(request adding.ValidatedRequest, genUUIDStr uuid.GenFunc) adding.PersistedCustomer {
		return adding.PersistedCustomer{
			CustomerId:  genUUIDStr(),
			FirstName:   adding.RetrieveFirstName(request.FirstName),
			LastName:    request.LastName,
			NationalId:  request.NationalId,
			PhoneNumber: request.PhoneNumber,
			AccountId:   request.AccountId,
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
