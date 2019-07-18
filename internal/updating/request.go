package updating

//Request represents the request to update a customer
type Request struct {
	CustomerID  string
	FirstName   string
	LastName    string
	NationalID  string
	PhoneNumber string
}
