package integration

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/albertowusuasare/customer-app/cmd/customer-svc/pkg"
	"github.com/albertowusuasare/customer-app/internal/api"
)

func TestRetreiveOneAPI(t *testing.T) {
	app := pkg.InmemApp()
	ts := httptest.NewServer(api.Handler(app))
	defer ts.Close()

	//create a customer
	existingCustomer := CreateTestDataCustomer(ts)
	customerID := existingCustomer.CustomerID

	retrievedCustomer := retrieveCustomer(ts, customerID)

	if customerID != retrievedCustomer.CustomerID {
		t.Fatalf("Expected customerID for retrieved customer %s got %s",
			customerID, retrievedCustomer.CustomerID)
	}

	testExpectedCustomerFieldsPresent(retrievedCustomer, t)

}

func retrieveCustomer(ts *httptest.Server, customerID string) api.CustomerRetrieveResponseDTO {
	customerRetrieveURL := ts.URL + "/customers/" + customerID
	res, _ := http.Get(customerRetrieveURL)
	response, _ := ioutil.ReadAll(res.Body)
	responseDTO := api.CustomerRetrieveResponseDTO{}
	UnMarshal(response, &responseDTO)
	return responseDTO
}

func testExpectedCustomerFieldsPresent(c api.CustomerRetrieveResponseDTO, t *testing.T) {
	var fields = []struct{ name, value string }{
		{
			"firstName",
			c.FirstName,
		},
		{
			"lastName",
			c.LastName,
		},
		{
			"nationalID",
			c.NationalID,
		},
		{
			"phoneNumber",
			c.PhoneNumber,
		},
		{
			"accountID",
			c.AccountID,
		},
		{
			"lastModifiedTime",
			c.LastModifiedTime,
		},
		{
			"createdTime",
			c.CreatedTime,
		},
		{
			"version",
			string(c.Version),
		},
	}

	for _, tt := range fields {
		t.Run(tt.name+"_present", func(t *testing.T) {
			if tt.value == "" {
				t.Errorf("got an empty value for field expecting non empty value")
			}
		})
	}
}
