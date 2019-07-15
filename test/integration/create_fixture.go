package integration

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/albertowusuasare/customer-app/internal/api"
)

// CreateTestDataCustomer returns a testdata customer
func CreateTestDataCustomer(ts *httptest.Server) api.CreateResponseDTO {
	requestBody, _ := ioutil.ReadFile("../data/create-request.json")
	buffer := bytes.NewBuffer(requestBody)
	res, _ := http.Post(ts.URL+"/customers/", "application/json", buffer)
	response, _ := ioutil.ReadAll(res.Body)
	responseDTO := api.CreateResponseDTO{}
	UnMarshal(response, &responseDTO)
	return responseDTO
}
