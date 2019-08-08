package app

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type requestInfo struct {
	HTTPMethod string   `json:"httpMethod"`
	URL        *url.URL `json:"httpRequestURI"`
}

func (i requestInfo) String() string {
	return fmt.Sprintf("httpMethod=%s, httpRequestURI=%s", i.HTTPMethod, i.URL.RequestURI())
}

// HTTPRequestLoggingFilter intercepts and logs data about incoming requests
type HTTPRequestLoggingFilter func() http.Handler

func (h HTTPRequestLoggingFilter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	i := requestInfo{HTTPMethod: r.Method, URL: r.URL}
	log.Printf("Request recieved: %+v", i)
	h().ServeHTTP(w, r)
}
