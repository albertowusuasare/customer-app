package updating

//UpdatedCustomer represents an updated customer in the datastore
type UpdatedCustomer struct {
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
