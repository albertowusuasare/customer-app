package updating

//UpdatedCustomer represents an updated customer in the datastore
type UpdatedCustomer struct {
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
