package workflow

import (
	"github.com/albertowusuasare/customer-app/internal/adding"
	"github.com/albertowusuasare/customer-app/internal/msg"
	"github.com/albertowusuasare/customer-app/internal/storage"
	"github.com/albertowusuasare/customer-app/internal/uuid"
)

// CreateFunc creates a customer from a request to add a new customer
type CreateFunc func(r adding.UnvalidatedRequest) (adding.Customer, error)

// Create is the default implementation of the customer create workflow
func Create(
	validateRequest adding.RequestValidatorFunc,
	genUUIDStr uuid.GenFunc,
	insertCustomer storage.InsertCustomerFunc,
	publishCustomerAdded msg.CustomerAddedPublisherFunc) CreateFunc {
	return func(request adding.UnvalidatedRequest) (adding.Customer, error) {
		validatedRequest, err := validateRequest(request)
		if err != nil {
			return adding.Customer{}, err

		}
		customer := insertCustomer(validatedRequest, genUUIDStr)
		publishCustomerAdded(customer)
		return customer, nil
	}
}
