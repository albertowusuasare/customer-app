package adding

// An UnPersistedCustomer is the value of a customer pre database persistence
type UnPersistedCustomer struct {
	FirstName   string
	LastName    string
	NationalId  string
	PhoneNumber string
	AccountId   string
}

// A PersistedCustomer is the value of a customer post database persistence
type PersistedCustomer struct {
	CustomerId  string
	FirstName   string
	LastName    string
	NationalId  string
	PhoneNumber string
	AccountId   string
}
