package inmem

import (
	"log"

	"github.com/albertowusuasare/customer-app/internal/adding"
	"github.com/albertowusuasare/customer-app/internal/msg"
	"github.com/albertowusuasare/customer-app/internal/updating"
)

// CustomerAddedPublisher returns an in memory customer add publisher
func CustomerAddedPublisher() msg.CustomerAddedPublisherFunc {
	return func(c *adding.Customer) msg.Response {
		publisher := func(messageID string) {
			log.Printf("Publishing customer add event with messageID=%+s", messageID)
		}
		return publish(*c, publisher)
	}
}

// CustomerUpdatedPublisher returns an in memory customer updated publisher
func CustomerUpdatedPublisher() msg.CustomerUpdatedPublisherFunc {
	return func(c updating.UpdatedCustomer) msg.Response {
		publisher := func(messageID string) {
			log.Printf("Publishing customer update event with messageID=%+s", messageID)
		}
		return publish(c, publisher)
	}
}

// CustomerRemovedPublisher returns an in memory customer removed publisher
func CustomerRemovedPublisher() msg.CustomerRemovedPublisherFunc {
	return func(ID string) msg.Response {
		publisher := func(messageID string) {
			log.Printf("Publishing customer removed event with messageID=%+s", messageID)
		}
		return publish(ID, publisher)
	}
}

type publisherFunc func(messageID string)

func publish(payload interface{}, doPublish publisherFunc) msg.Response {
	envelope := CreateEnvelope(payload)
	messageID := envelope.Header.ID
	doPublish(messageID)
	return msg.Response{
		MessageID:    messageID,
		Acknowledged: true,
	}
}
