package adding

import (
	"fmt"

	"github.com/albertowusuasare/customer-app/internal/uuid"
	"github.com/albertowusuasare/customer-app/internal/validation"
)

type firstName string
type lastName string
type customerID uuid.V4
type nationalID string
type phoneNumber string
type accountID uuid.V4

// A Customer is the value of a customer post insert
type Customer struct {
	CustomerID  customerID
	FirstName   firstName
	LastName    lastName
	NationalID  nationalID
	PhoneNumber phoneNumber
	AccountID   accountID
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

// RetrieveFirstName returns the underlying value for a firstName
func RetrieveFirstName(f firstName) string {
	return string(f)
}

// CreateLastName validates v and returns a lastname
// An error is returned if v is an invalid lastname
// TODO should this be unexported?
func CreateLastName(v string) (lastName, error) {
	if v == "" {
		return "", fmt.Errorf("LastName cannot be empty")
	}

	if !validation.IsUTFAlpahnumeric(v) {
		return "", fmt.Errorf("LastName must be alphanumeric")
	}

	if !validation.IsLengthLessOrEqual(v, 64) {
		return "", fmt.Errorf("LastName legnth must be less than or equal to 64")
	}

	return lastName(v), nil
}

// RetrieveLasttName returns the underlying value for a lastName
func RetrieveLasttName(v lastName) string {
	return string(v)
}

// CreateCustomerID and returns a customerID
func CreateCustomerID(v4UUID uuid.V4) customerID {
	return customerID(v4UUID)
}

// RetrieveCustomerID returns the underlying value for a customerID
func RetrieveCustomerID(v customerID) string {
	return string(v)
}

// TODO extra add validation rules
// CreateNationalID and returns a nationalID
// An error is returned if v is an invalid nationalID
func CreateNationalID(v string) (nationalID, error) {
	if v == "" {
		return "", fmt.Errorf("NationalID cannot be empty")
	}
	return nationalID(v), nil
}

// RetrieveNationalID returns the underlying value for a nationalID
func RetrieveNationalID(v nationalID) string {
	return string(v)
}

// TODO extra add validation rules
// CreatePhoneNumber and returns a phoneNumber
// An error is returned if v is an invalid phoneNumber
func CreatePhoneNumber(v string) (phoneNumber, error) {
	if v == "" {
		return "", fmt.Errorf("PhoneNumber cannot be empty")
	}
	return phoneNumber(v), nil
}

// RetrievePhoneNumber returns the underlying value for a phoneNumber
func RetrievePhoneNumber(v phoneNumber) string {
	return string(v)
}

// TODO extra add validation rules
// CreateAccountID and returns a phoneNumber
// An error is returned if v is an invalid accountID
func CreateAccountID(v string) (accountID, error) {
	if v == "" {
		return "", fmt.Errorf("AccountID cannot be empty")
	}

	if !uuid.IsValidUUID(v) {
		return "", fmt.Errorf("AccountID must be valid v4 UUID")
	}

	return accountID(uuid.V4(v)), nil
}

// RetrieveAccountID returns the underlying value for a nationalID
func RetrieveAccountID(v accountID) string {
	return string(v)
}
