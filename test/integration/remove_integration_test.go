package integration

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/albertowusuasare/customer-app/internal/app"
	"github.com/albertowusuasare/customer-app/internal/app/inmem"
)

func TestRemove(t *testing.T) {
	// Initialize test server
	inMemApp := inmem.App()
	ts := httptest.NewServer(app.Handler(inMemApp))

	// Create customer
	customer := CreateTestDataCustomer(ts)
	customerID := customer.CustomerID

	// Remove customer
	res := doRemove(customerID, ts)
	status := res.StatusCode
	expectedStatus := 204
	if status != expectedStatus {
		t.Errorf("Expected http status %d but got %d", expectedStatus, status)
	}

	// Assert customer does not exist
	testCustomerDoesNotExist(customerID, ts, t)
}

func doRemove(customerID string, ts *httptest.Server) *http.Response {
	URL := fmt.Sprintf("%s/customers/%s", ts.URL, customerID)
	req, reqErr := http.NewRequest(http.MethodDelete, URL, nil)
	if reqErr != nil {
		log.Fatal(reqErr)
	}
	client := &http.Client{}
	res, callErr := client.Do(req)
	if callErr != nil {
		log.Fatal(callErr)
	}
	return res
}

func testCustomerDoesNotExist(customerID string, ts *httptest.Server, t *testing.T) {
	URL := fmt.Sprintf("%s/customers/%s", ts.URL, customerID)
	res, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}
	statusCode := res.StatusCode
	if statusCode != http.StatusNotFound {
		t.Fatalf("customer retrieval for removed customer returned status %d instead of %d", statusCode, http.StatusNotFound)
	}

	t.Run("error_body_asertion", func(t *testing.T) {
		b, _ := ioutil.ReadAll(res.Body)
		errDTO := app.CustomerRetrieveErrorDTO{}
		UnMarshal(b, &errDTO)

		e := fmt.Sprintf("No record exits for customerID=%s", customerID)
		a := errDTO.Message
		if e != a {
			t.Errorf("Expected error message=%s but got=%s", e, a)
		}
	})

}
