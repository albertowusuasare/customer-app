package handler

import (
	"net/http"
	"strings"
)

func RetrieveCustomerId(r *http.Request) string {
	return strings.Split(r.RequestURI, "/")[2]
}
