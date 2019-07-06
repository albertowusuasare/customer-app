package adding

type Request struct {
	FirstName   string
	LastName    string
	NationalId  string
	PhoneNumber string
	AccountId   string
}

type RequestValidatorFunc func(r Request) (UnPersistedCustomer, error)

func RequestToUnPersistedCustomer(r Request) (UnPersistedCustomer, error) {
	return UnPersistedCustomer{
		FirstName:   r.FirstName,
		LastName:    r.LastName,
		NationalId:  r.NationalId,
		PhoneNumber: r.PhoneNumber,
		AccountId:   r.AccountId,
	}, nil
}
