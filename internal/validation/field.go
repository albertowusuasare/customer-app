package validation

import "fmt"

// FieldName is the name of the field that failed validation
type FieldName string

// Message is the message describing the field validation failure
type Message string

// Fields is a key value collection of the names of all failed fields
// and their corresponding failure messages
type Fields map[FieldName]Message

// Implemententors of FieldValidation adds a means to get validation error fields from an error
type fieldValidationFailure interface {
	FailedFields() Fields
}

// Error is a representation of all input validation errors when initializing a workflow
type Error struct {
	Fields map[FieldName]Message
}

// FailedFields returns the fields that failed validation
func (v Error) FailedFields() Fields {
	result := map[FieldName]Message{}
	for k, v := range v.Fields {
		result[FieldName(k)] = Message(v)
	}
	return result
}

// Error adds the error behavior to ValidationError
func (v Error) Error() string {
	return fmt.Sprintf("Validation error has fields %s", v.Fields)

}

// IsFieldValidationError returns trye if the error is of type fieldValidationFailure
func IsFieldValidationError(err error) bool {
	_, ok := err.(fieldValidationFailure)
	return ok
}

// GetFailedValidationFields returns the failed validation fields or error if the error
// is not a fieldValidationFailure
func GetFailedValidationFields(err error) (map[FieldName]Message, error) {
	ve, ok := err.(fieldValidationFailure)

	if ok {
		return ve.FailedFields(), nil
	}

	return map[FieldName]Message{},
		fmt.Errorf("Unable to obtain failued validation fields from non field validation error %+v", err)
}

// RetrieveFieldName returns the underlying field name for f
func RetrieveFieldName(f FieldName) string {
	return string(f)
}

// RetrieveMessage returns the underlying message for m
func RetrieveMessage(m Message) string {
	return string(m)
}
