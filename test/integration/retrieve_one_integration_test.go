package integration

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/albertowusuasare/customer-app/internal/app"
	"github.com/albertowusuasare/customer-app/internal/app/inmem"
)

func TestRetreiveOneAPI(t *testing.T) {
	inmemApp := inmem.App()
	ts := httptest.NewServer(app.Handler(inmemApp))
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

func retrieveCustomer(ts *httptest.Server, customerID string) app.CustomerRetrieveResponseDTO {
	customerRetrieveURL := ts.URL + "/customers/" + customerID
	res, _ := http.Get(customerRetrieveURL)
	response, _ := ioutil.ReadAll(res.Body)
	responseDTO := app.CustomerRetrieveResponseDTO{}
	UnMarshal(response, &responseDTO)
	return responseDTO
}

func testExpectedCustomerFieldsPresent(c app.CustomerRetrieveResponseDTO, t *testing.T) {
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
