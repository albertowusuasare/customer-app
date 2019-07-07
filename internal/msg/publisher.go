package msg

import (
	"github.com/albertowusuasare/customer-app/internal/adding"
	"github.com/albertowusuasare/customer-app/internal/updating"
)

// Response represents the response after a message is published
type Response struct {
	MessageId    string
	Acknowledged bool
}

// CustomerAddedPublisherFunc publishes an event for the addition of a persisted customer
type CustomerAddedPublisherFunc func(customer adding.PersistedCustomer) Response

// CustomerUpdatedPublisherFunc publishes an event for an update to a persisted customer
type CustomerUpdatedPublisherFunc func(customer updating.UpdatedCustomer) Response

// CustomerRemovedPublisherFunc publishes an event for the removal of a persisted customer
type CustomerRemovedPublisherFunc func(customerId string) Response
