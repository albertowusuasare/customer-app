package adding

import (
	"fmt"

	"github.com/albertowusuasare/customer-app/internal/uuid"
	"github.com/albertowusuasare/customer-app/internal/validation"
)

// FirstName is a representation of a customer's firstname
type FirstName string

// LastName is a representation of a customer's lastname
type LastName string

// CustomerID is a uuid v4 represenation of a customer's identifier
type CustomerID uuid.V4

// NationalID is a representation of customer's national identifier
type NationalID string

// PhoneNumber is a representation of a customer's phone number
type PhoneNumber string

// AccountID is a uuid v4 representation of a customer's account identifier
type AccountID uuid.V4

// A Customer is the value of a customer post insert
type Customer struct {
	id          CustomerID
	firstName   FirstName
	lastName    LastName
	nationalID  NationalID
	phoneNumber PhoneNumber
	accountID   AccountID
}

// NewCustomer creates a new customer for a validated request
func NewCustomer(r ValidatedRequest, id uuid.V4) *Customer {
	return &Customer{
		id:          CustomerID(id),
		firstName:   r.firstName,
		lastName:    r.lastName,
		nationalID:  r.nationalID,
		phoneNumber: r.phoneNumber,
		accountID:   r.accountID,
	}
}

// RetrieveFirstName returns the customer's firstname
func (c *Customer) RetrieveFirstName() FirstName {
	return c.firstName
}

// RetrieveLastName returns the customer's lastName
func (c *Customer) RetrieveLastName() LastName {
	return c.lastName
}

// RetrieveCustomerID returns the customer's ID
func (c *Customer) RetrieveCustomerID() CustomerID {
	return c.id
}

// RetrieveNationalID returns the customer's national identifier
func (c *Customer) RetrieveNationalID() NationalID {
	return c.nationalID
}

// RetrievePhoneNumber returns the customer's  phoneNumber
func (c *Customer) RetrievePhoneNumber() PhoneNumber {
	return c.phoneNumber
}

// RetrieveAccountID returns the customer's account identifier
func (c *Customer) RetrieveAccountID() AccountID {
	return c.accountID
}

// CreateFirstName validates v and returns a firstname
// An error is returned if v is an invalid firstname
func CreateFirstName(v string) (FirstName, error) {
	if v == "" {
		return "", fmt.Errorf("Firstname cannot be empty")
	}

	if !validation.IsUTFAlpahnumeric(v) {
		return "", fmt.Errorf("Firstname must be alphanumeric")
	}

	if !validation.IsLengthLessOrEqual(v, 64) {
		return "", fmt.Errorf("FirstName legnth must be less than or equal to 64")
	}

	return FirstName(v), nil
}

// CreateLastName validates v and returns a lastname
// An error is returned if v is an invalid lastname
// TODO should this be unexported?
func CreateLastName(v string) (LastName, error) {
	if v == "" {
		return "", fmt.Errorf("LastName cannot be empty")
	}

	if !validation.IsUTFAlpahnumeric(v) {
		return "", fmt.Errorf("LastName must be alphanumeric")
	}

	if !validation.IsLengthLessOrEqual(v, 64) {
		return "", fmt.Errorf("LastName legnth must be less than or equal to 64")
	}

	return LastName(v), nil
}

// CreateCustomerID and returns a customerID
func CreateCustomerID(v4UUID uuid.V4) CustomerID {
	return CustomerID(v4UUID)
}

// TODO extra add validation rules
// CreateNationalID and returns a nationalID
// An error is returned if v is an invalid nationalID
func CreateNationalID(v string) (NationalID, error) {
	if v == "" {
		return "", fmt.Errorf("NationalID cannot be empty")
	}
	return NationalID(v), nil
}

// TODO extra add validation rules
// CreatePhoneNumber and returns a phoneNumber
// An error is returned if v is an invalid phoneNumber
func CreatePhoneNumber(v string) (PhoneNumber, error) {
	if v == "" {
		return "", fmt.Errorf("PhoneNumber cannot be empty")
	}
	return PhoneNumber(v), nil
}

// TODO extra add validation rules
// CreateAccountID and returns a phoneNumber
// An error is returned if v is an invalid accountID
func CreateAccountID(v string) (AccountID, error) {
	if v == "" {
		return "", fmt.Errorf("AccountID cannot be empty")
	}

	if !uuid.IsValidUUID(v) {
		return "", fmt.Errorf("AccountID must be valid v4 UUID")
	}

	return AccountID(uuid.V4(v)), nil
}
