package msg

import (
	"github.com/albertowusuasare/customer-app/internal/adding"
	"github.com/albertowusuasare/customer-app/internal/updating"
)

type Response struct {
	MessageId    string
	Acknowledged bool
}

type CustomerAddedPublisherFunc func(customer adding.PersistedCustomer) Response

type CustomerUpdatedPublisherFunc func(customer updating.UpdatedCustomer) Response

type CustomerRemovedPublisherFunc func(customerId string) Response
