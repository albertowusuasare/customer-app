package validation

import "unicode"

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
