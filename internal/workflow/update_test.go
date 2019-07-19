package workflow

import (
	"testing"

	"github.com/albertowusuasare/customer-app/internal/msg"
	"github.com/albertowusuasare/customer-app/internal/storage"
	"github.com/albertowusuasare/customer-app/internal/updating"
)

func TestUpdate(t *testing.T) {
	request := mockUpdateRequest()
	expected := mockExpectedUpdatedCustomer(request)
	updateDb := updateDb(request, expected, t)
	publishUpdate := publishUpdate(expected, t)

	update := Update(updateDb, publishUpdate)
	actual := update(request)

	if expected != actual {
		t.Errorf("expectedUpdateResponse=%+v is not equal to actualUpdateResponse=%+v", expected, actual)
	}
}

func mockUpdateRequest() updating.Request {
	return updating.Request{
		CustomerID:  "415eb201-83ed-48b6-b26c-30271a492e4b",
		FirstName:   "John",
		LastName:    "Doe",
		NationalID:  "9876543",
		PhoneNumber: "987654321",
	}
}

func mockExpectedUpdatedCustomer(r updating.Request) updating.UpdatedCustomer {
	return updating.UpdatedCustomer{
		CustomerID:       r.CustomerID,
		FirstName:        r.FirstName,
		LastName:         r.LastName,
		NationalID:       r.NationalID,
		PhoneNumber:      r.PhoneNumber,
		AccountID:        "1bcf7c0c-9174-477e-a490-cac4e42af8f2",
		LastModifiedTime: "2019-07-05T01:39:20+01:00",
		CreatedTime:      "2019-07-05T01:39:20+01:00",
		Version:          1,
	}
}

func updateDb(
	expectedRequest updating.Request,
	expectedUpdatedCus updating.UpdatedCustomer,
	t *testing.T) storage.UpdateCustomerFunc {
	return func(request updating.Request) updating.UpdatedCustomer {
		if expectedRequest != request {
			t.Errorf("expectedRequest=%+v is not equal to actualRequest=%+v", expectedRequest, request)
		}
		return expectedUpdatedCus
	}
}

func publishUpdate(
	expectedUpdatedCus updating.UpdatedCustomer,
	t *testing.T) msg.CustomerUpdatedPublisherFunc {
	return func(customer updating.UpdatedCustomer) msg.Response {
		if expectedUpdatedCus != customer {
			t.Errorf("expectedPublishArgr=%+v is not equal to actualPublishArg=%+v", expectedUpdatedCus, customer)
		}
		return msg.Response{
			MessageID:    "3e1c64d3-9bb6-40f6-8acf-9e38c6acd6cd",
			Acknowledged: true,
		}
	}
}
