package adding

import (
	"fmt"

	"github.com/albertowusuasare/customer-app/internal/validation"
)

type firstName string

// A PersistedCustomer is the value of a customer post database persistence
type PersistedCustomer struct {
	CustomerId  string
	FirstName   string
	LastName    string
	NationalId  string
	PhoneNumber string
	AccountId   string
}

/* Simple type constructors */

// CreateFirstName validates v and returns a firstname
// An error is returned if v is an invalid firstname
func CreateFirstName(v string) (firstName, error) {
	if v == "" {
		return "", fmt.Errorf("Firstname cannot be empty")
	}

	if !validation.IsUTFAlpahnumeric(v) {
		return "", fmt.Errorf("Firstname must be alphanumeric")
	}

	if !validation.IsLengthLessOrEqual(v, 64) {
		return "", fmt.Errorf("FirstName legnth must be less than or equal to 64")
	}

	return firstName(v), nil
}

// RetrieveFirstName the underlying value for a firstName
func RetrieveFirstName(f firstName) string {
	return string(f)
}
