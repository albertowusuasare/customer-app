package integration

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/albertowusuasare/customer-app/cmd/customer-svc/handler"
)

// CreateTestDataCustomer returns a testdata customer
func CreateTestDataCustomer(ts *httptest.Server) handler.CreateResponseDTO {
	requestBody, _ := ioutil.ReadFile("../data/create-request.json")
	buffer := bytes.NewBuffer(requestBody)
	res, _ := http.Post(ts.URL+"/customers/", "application/json", buffer)
	response, _ := ioutil.ReadAll(res.Body)
	responseDTO := handler.CreateResponseDTO{}
	UnMarshal(response, &responseDTO)
	return responseDTO
}
