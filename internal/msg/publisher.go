package msg

import (
	"log"

	"github.com/albertowusuasare/customer-app/internal/adding"
	"github.com/albertowusuasare/customer-app/internal/updating"
)

// Response represents the response after a message is published
type Response struct {
	MessageID    string
	Acknowledged bool
}

// CustomerAddedPublisherFunc publishes an event for the addition of a customer
type CustomerAddedPublisherFunc func(customer *adding.Customer) Response

// CustomerUpdatedPublisherFunc publishes an event for an update to a customer
type CustomerUpdatedPublisherFunc func(customer updating.UpdatedCustomer) Response

// CustomerRemovedPublisherFunc publishes an event for the removal of a customer
type CustomerRemovedPublisherFunc func(customerId string) Response

// LoggingCustomerAddedPublisher is a CustomerAddedPublisherFunc that logs the publisher response
func LoggingCustomerAddedPublisher(delegate CustomerAddedPublisherFunc) CustomerAddedPublisherFunc {
	return func(c *adding.Customer) Response {
		r := delegate(c)
		log.Printf("Customer add message published. messageID=%s", r.MessageID)
		return r
	}
}

// LoggingCustomerUpdatedPublisher is a CustomerUpdatedPublisherFunc that logs the publisher response
func LoggingCustomerUpdatedPublisher(delegate CustomerUpdatedPublisherFunc) CustomerUpdatedPublisherFunc {
	return func(c updating.UpdatedCustomer) Response {
		r := delegate(c)
		log.Printf("Customer update message published. messageID=%s", r.MessageID)
		return r
	}
}

// LoggingCustomerRemovedPublisher is a CustomerRemovedPublisherFunc that logs the publisher response
func LoggingCustomerRemovedPublisher(delegate CustomerRemovedPublisherFunc) CustomerRemovedPublisherFunc {
	return func(ID string) Response {
		r := delegate(ID)
		log.Printf("Customer remove message published. messageID=%s", r.MessageID)
		return r
	}
}
