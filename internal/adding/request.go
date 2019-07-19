package adding

import "github.com/albertowusuasare/customer-app/internal/validation"

// An UnvalidatedRequest is the unvalidated request to add a new customer.
type UnvalidatedRequest struct {
	FirstName   string
	LastName    string
	NationalID  string
	PhoneNumber string
	AccountID   string
}

// A ValidatedRequest is the value of a customer add request post validation
type ValidatedRequest struct {
	FirstName   firstName
	LastName    lastName
	NationalID  string
	PhoneNumber string
	AccountID   string
}

// A RequestValidatorFunc exposes functionaltiy to validate an incoming add request.
type RequestValidatorFunc func(r UnvalidatedRequest) (ValidatedRequest, error)

// ValidateRequest is the primary validator for incoming customer add requests.
func ValidateRequest(r UnvalidatedRequest) (ValidatedRequest, error) {
	failedFields := map[validation.FieldName]validation.Message{}
	firstName, firstNameErr := CreateFirstName(r.FirstName)
	lastName, lastNameErr := CreateLastName(r.LastName)

	if firstNameErr != nil {
		failedFields[validation.FieldName("firstName")] = validation.Message(firstNameErr.Error())
	}

	if lastNameErr != nil {
		failedFields[validation.FieldName("lastName")] = validation.Message(lastNameErr.Error())
	}

	if len(failedFields) != 0 {
		validationError := validation.Error{Fields: failedFields}
		return ValidatedRequest{}, validationError
	}

	return ValidatedRequest{
		FirstName:   firstName,
		LastName:    lastName,
		NationalID:  r.NationalID,
		PhoneNumber: r.PhoneNumber,
		AccountID:   r.AccountID,
	}, nil
}
