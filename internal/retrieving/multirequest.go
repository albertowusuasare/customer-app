package retrieving

// MultiRequest is the definition of the query from which customers will be fetched.
// Key Value pairs of query field and query value are captured in QueryParams
// e.g key="FirstName" value="John", key ="LastName" value="Doe"
type MultiRequest struct {
	QueryParams map[string]string
}
