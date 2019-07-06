package adding

type UnPersistedCustomer struct {
	FirstName   string
	LastName    string
	NationalId  string
	PhoneNumber string
	AccountId   string
}

type PersistedCustomer struct {
	CustomerId  string
	FirstName   string
	LastName    string
	NationalId  string
	PhoneNumber string
	AccountId   string
}
