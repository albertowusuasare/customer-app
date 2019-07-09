package retrieving

import "fmt"

// CustomerNonExistent is a representation of the error returned when a non existent customer is retrieved
type CustomerNonExistent struct {
	CustomerID string
}

// Error is the implementation of the Error interface on CustomerNonExistent
func (e CustomerNonExistent) Error() string {
	return fmt.Sprintf("No record exits for customerID=%s", e.CustomerID)
}
