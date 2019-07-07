package updating

//Request represents the request to update a customer
type Request struct {
	CustomerId  string
	FirstName   string
	LastName    string
	NationalId  string
	PhoneNumber string
}
