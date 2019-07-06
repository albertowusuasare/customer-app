package kafka

import (
	"customer-app/internal/adding"
	"customer-app/internal/msg"
	"customer-app/internal/updating"
)

func CustomerAddedPublisher() msg.CustomerAddedPublisherFunc {
	return func(c adding.PersistedCustomer) msg.Response {
		panic("Customer added kafka publisher not yet implemented")
	}
}

func CustomerUpdatedPublisher() msg.CustomerUpdatedPublisherFunc {
	return func(c updating.UpdatedCustomer) msg.Response {
		panic("Customer updated kafka publisher not yet implemented")
	}
}

func CustomerRemovedPublisher() msg.CustomerRemovedPublisherFunc {
	return func(id string) msg.Response {
		panic("Customer removed kafka publisher not yet implemented")
	}
}
