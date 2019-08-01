package workflow

import (
	"github.com/albertowusuasare/customer-app/internal/adding"
	"github.com/albertowusuasare/customer-app/internal/msg"
	"github.com/albertowusuasare/customer-app/internal/storage"
	"github.com/albertowusuasare/customer-app/internal/uuid"
	"github.com/pkg/errors"
)

// CreateFunc creates a customer from a request to add a new customer
type CreateFunc func(r adding.UnvalidatedRequest) (*adding.Customer, error)

// Create is the default implementation of the customer create workflow
func Create(
	validateRequest adding.RequestValidatorFunc,
	genV4UUID uuid.GenV4Func,
	insertCustomer storage.InsertCustomerFunc,
	publishCustomerAdded msg.CustomerAddedPublisherFunc) CreateFunc {
	return func(request adding.UnvalidatedRequest) (*adding.Customer, error) {
		validatedRequest, err := validateRequest(request)
		if err != nil {
			return nil, errors.Wrap(err, "Customer create validation failure")
		}
		customer := adding.NewCustomer(validatedRequest, genV4UUID())

		insertErr := insertCustomer(customer)
		if insertErr != nil {
			return nil, errors.Wrap(err, "Failure during customer insert")
		}

		publishCustomerAdded(customer)
		return customer, nil
	}
}
