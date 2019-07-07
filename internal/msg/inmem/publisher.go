package inmem

import (
	"log"

	"github.com/albertowusuasare/customer-app/internal/adding"
	"github.com/albertowusuasare/customer-app/internal/msg"
	"github.com/albertowusuasare/customer-app/internal/updating"
)

// CustomerAddedPublisher returns an in memory customer add publisher
func CustomerAddedPublisher() msg.CustomerAddedPublisherFunc {
	return func(c adding.PersistedCustomer) msg.Response {
		publisher := func(messageId string) {
			log.Printf("Publishing customer add event with messageId=%+s", messageId)
		}
		return publish(c, publisher)
	}
}

// CustomerUpdatedPublisher returns an in memory customer updated publisher
func CustomerUpdatedPublisher() msg.CustomerUpdatedPublisherFunc {
	return func(c updating.UpdatedCustomer) msg.Response {
		publisher := func(messageId string) {
			log.Printf("Publishing customer update event with messageId=%+s", messageId)
		}
		return publish(c, publisher)
	}
}

// CustomerRemovedPublisher returns an in memory customer removed publisher
func CustomerRemovedPublisher() msg.CustomerRemovedPublisherFunc {
	return func(id string) msg.Response {
		publisher := func(messageId string) {
			log.Printf("Publishing customer removed event with messageId=%+s", messageId)
		}
		return publish(id, publisher)
	}
}

type publisherFunc func(messageId string)

func publish(payload interface{}, doPublish publisherFunc) msg.Response {
	envelope := CreateEnvelope(payload)
	messageId := envelope.Header.Id
	doPublish(messageId)
	return msg.Response{
		MessageId:    messageId,
		Acknowledged: true,
	}
}
