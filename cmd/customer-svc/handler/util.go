package handler

import (
	"net/http"
	"strings"
)

// RetrieveCustomerId retrieves the customerId from an http request
func RetrieveCustomerId(r *http.Request) string {
	return strings.Split(r.RequestURI, "/")[2]
}
