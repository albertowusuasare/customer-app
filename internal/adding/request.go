package adding

import "github.com/albertowusuasare/customer-app/internal/validation"

// An UnvalidatedRequest is the unvalidated request to add a new customer.
type UnvalidatedRequest struct {
	FirstName   string
	LastName    string
	NationalId  string
	PhoneNumber string
	AccountId   string
}

// A ValidatedRequest is the value of a customer add request post validation
type ValidatedRequest struct {
	FirstName   firstName
	LastName    string
	NationalId  string
	PhoneNumber string
	AccountId   string
}

// A RequestValidatorFunc exposes functionaltiy to validate an incoming add request.
type RequestValidatorFunc func(r UnvalidatedRequest) (ValidatedRequest, error)

// ValidateRequest is the primary validator for incoming customer add requests.
func ValidateRequest(r UnvalidatedRequest) (ValidatedRequest, error) {
	failedFields := map[validation.FieldName]validation.Message{}
	firstName, err := CreateFirstName(r.FirstName)

	if err != nil {
		failedFields[validation.FieldName("firstName")] = validation.Message(err.Error())
	}

	if len(failedFields) != 0 {
		validationError := validation.Error{Fields: failedFields}
		return ValidatedRequest{}, validationError
	}

	return ValidatedRequest{
		FirstName:   firstName,
		LastName:    r.LastName,
		NationalId:  r.NationalId,
		PhoneNumber: r.PhoneNumber,
		AccountId:   r.AccountId,
	}, nil
}
