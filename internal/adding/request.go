package adding

// A Request is the unvalidated request to add a new customer.
// A Request will need to be validated in order to consider it being persisted in the database
type Request struct {
	FirstName   string
	LastName    string
	NationalId  string
	PhoneNumber string
	AccountId   string
}

// A RequestValidatorFunc exposes functionaltiy to validate an incoming add request.
// The output of the validator is the UnPersistedCustomer to be saved if no errors are encountered.
type RequestValidatorFunc func(r Request) (UnPersistedCustomer, error)

// RequestToUnPersistedCustomer is the primary validator for incoming customer add requests.
func RequestToUnPersistedCustomer(r Request) (UnPersistedCustomer, error) {
	return UnPersistedCustomer(r), nil
}
