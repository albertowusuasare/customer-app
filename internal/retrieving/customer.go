package retrieving

// A Customer represents all the data for a persisted customer
// The data includes audit data such as CreatedTime, lastModifiedTime, Version
type Customer struct {
	CustomerId       string
	FirstName        string
	LastName         string
	NationalId       string
	PhoneNumber      string
	AccountId        string
	LastModifiedTime string
	CreatedTime      string
	Version          int
}
