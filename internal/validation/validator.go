package validation

import (
	"unicode"

	"github.com/google/uuid"
)

// IsUTFAlpahnumeric checks if the string contains only unicode letters and numbers. Empty string is valid.
func IsUTFAlpahnumeric(s string) bool {
	if s == "" {
		return true
	}
	chars := []rune(s)
	f := chars[0]
	return (unicode.IsLetter(f) || unicode.IsNumber(f)) &&
		IsUTFAlpahnumeric(string(chars[1:]))
}

// IsLengthLessOrEqual checks if the string has length less than max
func IsLengthLessOrEqual(s string, max int) bool {
	return len(s) <= max
}

// IsValidUUID checks if a given string is a valid v4 UUID
func IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}
