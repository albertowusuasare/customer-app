package inmem

import (
	"customer-app/internal/adding"
	"customer-app/internal/msg"
	"customer-app/internal/updating"
	"log"
)

func CustomerAddedPublisher() msg.CustomerAddedPublisherFunc {
	return func(c adding.PersistedCustomer) msg.Response {
		publisher := func(messageId string) {
			log.Printf("Publishing customer add event with messageId=%+s", messageId)
		}
		return publish(c, publisher)
	}
}

func CustomerUpdatedPublisher() msg.CustomerUpdatedPublisherFunc {
	return func(c updating.UpdatedCustomer) msg.Response {
		publisher := func(messageId string) {
			log.Printf("Publishing customer update event with messageId=%+s", messageId)
		}
		return publish(c, publisher)
	}
}

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
