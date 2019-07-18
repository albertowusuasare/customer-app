package kafka

import (
	"github.com/albertowusuasare/customer-app/internal/adding"
	"github.com/albertowusuasare/customer-app/internal/msg"
	"github.com/albertowusuasare/customer-app/internal/updating"
)

// CustomerAddedPublisher returns a kafka customer add publisher
func CustomerAddedPublisher() msg.CustomerAddedPublisherFunc {
	return func(c adding.Customer) msg.Response {
		panic("Customer added kafka publisher not yet implemented")
	}
}

// CustomerUpdatedPublisher returns a kafka customer updated publisher
func CustomerUpdatedPublisher() msg.CustomerUpdatedPublisherFunc {
	return func(c updating.UpdatedCustomer) msg.Response {
		panic("Customer updated kafka publisher not yet implemented")
	}
}

// CustomerRemovedPublisher returns a kafka customer removed publisher
func CustomerRemovedPublisher() msg.CustomerRemovedPublisherFunc {
	return func(id string) msg.Response {
		panic("Customer removed kafka publisher not yet implemented")
	}
}
