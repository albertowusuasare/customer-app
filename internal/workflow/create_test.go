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
	expectedCustomer := newCustomer(expectedID, request)

	requestValidator := successRequestValidator()
	genUUIDStr := func() string { return expectedID }
	insertCustomer := successInsertCustomer()
	publishCustomerAdded := successCustomerAddedPublisher(expectedCustomer, t)

	createFunc := Create(requestValidator, genUUIDStr, insertCustomer, publishCustomerAdded)
	actualCustomer, _ := createFunc(request)

	if expectedCustomer != actualCustomer {
		t.Errorf("expectedCustomer=%+v is not equal to actualCustomer=%+v", expectedCustomer, actualCustomer)
	}

}

func mockAddingRequest() adding.UnvalidatedRequest {
	return adding.UnvalidatedRequest{
		FirstName:   "John",
		LastName:    "Doe",
		NationalID:  "987654321",
		PhoneNumber: "020765432",
		AccountID:   "b253fd2a-6bb9-49db-9fb5-f6388e1661a7",
	}
}

func newCustomer(ID string, r adding.UnvalidatedRequest) adding.Customer {
	return adding.Customer{
		CustomerID:  ID,
		FirstName:   r.FirstName,
		LastName:    r.LastName,
		NationalID:  r.NationalID,
		PhoneNumber: r.PhoneNumber,
		AccountID:   r.AccountID,
	}
}

func successRequestValidator() adding.RequestValidatorFunc {
	return func(r adding.UnvalidatedRequest) (adding.ValidatedRequest, error) {
		firstName, _ := adding.CreateFirstName(r.FirstName)
		return adding.ValidatedRequest{
			FirstName:   firstName,
			LastName:    r.LastName,
			NationalID:  r.NationalID,
			PhoneNumber: r.PhoneNumber,
			AccountID:   r.AccountID,
		}, nil
	}
}

func successInsertCustomer() storage.InsertCustomerFunc {
	return func(request adding.ValidatedRequest, genUUIDStr uuid.GenFunc) adding.Customer {
		return adding.Customer{
			CustomerID:  genUUIDStr(),
			FirstName:   adding.RetrieveFirstName(request.FirstName),
			LastName:    request.LastName,
			NationalID:  request.NationalID,
			PhoneNumber: request.PhoneNumber,
			AccountID:   request.AccountID,
		}
	}
}

func successCustomerAddedPublisher(expectedCustomer adding.Customer,
	t *testing.T) msg.CustomerAddedPublisherFunc {
	return func(customer adding.Customer) msg.Response {
		if expectedCustomer != customer {
			t.Errorf(fmt.Sprintf("Invalid customerAddedPublish arg. Expected=%+v Actual=%+v", expectedCustomer, customer))
		}
		return msg.Response{
			MessageID:    "19a1eed3-a650-412d-aeb7-20fabe0b37bc",
			Acknowledged: true,
		}
	}

}
