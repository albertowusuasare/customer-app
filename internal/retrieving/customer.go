package retrieving

// A Customer represents all the data for a  customer
// The data includes audit data such as CreatedTime, lastModifiedTime, Version
type Customer struct {
	CustomerID       string
	FirstName        string
	LastName         string
	NationalID       string
	PhoneNumber      string
	AccountID        string
	LastModifiedTime string
	CreatedTime      string
	Version          int
}
