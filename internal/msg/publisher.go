package msg

import (
	"github.com/albertowusuasare/customer-app/internal/adding"
	"github.com/albertowusuasare/customer-app/internal/updating"
)

// Response represents the response after a message is published
type Response struct {
	MessageID    string
	Acknowledged bool
}

// CustomerAddedPublisherFunc publishes an event for the addition of a customer
type CustomerAddedPublisherFunc func(customer adding.Customer) Response

// CustomerUpdatedPublisherFunc publishes an event for an update to a customer
type CustomerUpdatedPublisherFunc func(customer updating.UpdatedCustomer) Response

// CustomerRemovedPublisherFunc publishes an event for the removal of a customer
type CustomerRemovedPublisherFunc func(customerId string) Response
