package workflow

import (
	"testing"

	"github.com/albertowusuasare/customer-app/internal/msg"
	"github.com/albertowusuasare/customer-app/internal/storage"
)

func TestRemove(t *testing.T) {
	expectedId := "65f3bf6e-cf8f-4958-8181-353cd21e794c"

	removeFromDb := removeDbCustomer(expectedId, t)
	publishRemove := publishRemove(expectedId, t)
	remove := Remove(removeFromDb, publishRemove)

	remove(expectedId)
}

func removeDbCustomer(expectedId string, t *testing.T) storage.RemoveCustomerFunc {
	return func(id string) {
		if expectedId != id {
			t.Errorf("removeDbCustomer invoked with customerId=%s, intead of customerId=%s",
				id, expectedId)
		}
	}
}

func publishRemove(expectedId string, t *testing.T) msg.CustomerRemovedPublisherFunc {
	return func(id string) msg.Response {
		if expectedId != id {
			t.Errorf("publishRemove invoked with customerId=%s, intead of customerId=%s",
				id, expectedId)
		}
		return msg.Response{
			MessageId:    "2c0bec20-a729-4ed3-b343-ea132297a1a3",
			Acknowledged: true,
		}
	}
}
