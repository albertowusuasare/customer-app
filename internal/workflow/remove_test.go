package workflow

import (
	"testing"

	"github.com/albertowusuasare/customer-app/internal/msg"
	"github.com/albertowusuasare/customer-app/internal/storage"
)

func TestRemove(t *testing.T) {
	expectedID := "65f3bf6e-cf8f-4958-8181-353cd21e794c"

	removeFromDb := removeDbCustomer(expectedID, t)
	publishRemove := publishRemove(expectedID, t)
	remove := Remove(removeFromDb, publishRemove)

	remove(expectedID)
}

func removeDbCustomer(expectedID string, t *testing.T) storage.RemoveCustomerFunc {
	return func(ID string) {
		if expectedID != ID {
			t.Errorf("removeDbCustomer invoked with customerID=%s, intead of customerID=%s",
				ID, expectedID)
		}
	}
}

func publishRemove(expectedID string, t *testing.T) msg.CustomerRemovedPublisherFunc {
	return func(ID string) msg.Response {
		if expectedID != ID {
			t.Errorf("publishRemove invoked with customerID=%s, intead of customerID=%s",
				ID, expectedID)
		}
		return msg.Response{
			MessageID:    "2c0bec20-a729-4ed3-b343-ea132297a1a3",
			Acknowledged: true,
		}
	}
}
