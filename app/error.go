package app

// ErrorCode is a machine friendly description of the error
type ErrorCode string

// ErrorMessage is a human friendly description of the error
type ErrorMessage string

// ErrorParams captures an additional key value pairs of information about an error
type ErrorParams map[string]string

// An Error is a representation of an api error
// Errors include but not limited to domain errors, validation, authentication error etc
type Error struct {
	Code    ErrorCode    `json:"code"`
	Message ErrorMessage `json:"message"`
	Params  ErrorParams  `json:"params"`
}
