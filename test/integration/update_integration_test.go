package integration

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/albertowusuasare/customer-app/app"
	"github.com/albertowusuasare/customer-app/app/inmem"
)

func TestUpdate(t *testing.T) {
	inMemApp := inmem.App()
	ts := httptest.NewServer(app.Handler(inMemApp))

	// Seed customer
	customer := CreateTestDataCustomer(ts)

	// Create update request
	testdata := "../data/update-request.json"
	req, err := ioutil.ReadFile(testdata)
	if err != nil {
		log.Fatalf("Unable to read update request test data from %s", testdata)
	}
	// Update customer
	URL := fmt.Sprintf("%s/customers/%s", ts.URL, customer.CustomerID)
	client := &http.Client{}
	request, _ := http.NewRequest(http.MethodPut, URL, bytes.NewBuffer(req))
	res, _ := client.Do(request)
	responseBody, _ := ioutil.ReadAll(res.Body)

	// Assert api response
	if res.Status != "200 OK" {
		t.Fatalf("Update request failure expecting:200 OK got:%s", res.Status)
	}

	testUpdateResponsePayload(req, responseBody, t)
}

func testUpdateResponsePayload(req []byte, res []byte, t *testing.T) {

	requestDTO := app.UpdateRequestDTO{}
	UnMarshal(req, &requestDTO)

	responseDTO := app.UpdateResponseDTO{}
	UnMarshal(res, &responseDTO)

	log.Printf("Response %+v", responseDTO)

	//test response fields matches request
	requestFieldsTests := []struct {
		name,
		expected,
		actual string
	}{
		{"firstName", requestDTO.FirstName, responseDTO.FirstName},
		{"lastName", requestDTO.LastName, responseDTO.LastName},
		{"nationalId", requestDTO.NationalID, responseDTO.NationalID},
		{"PhoneNumber", requestDTO.PhoneNumber, responseDTO.PhoneNumber},
		{"Version", "1", strconv.Itoa(responseDTO.Version)},
	}

	for _, tt := range requestFieldsTests {
		t.Run(tt.name+"_match", func(t *testing.T) {
			a := tt.actual
			e := tt.expected
			if a != e {
				t.Errorf("got %s want %s", a, e)
			}
		})
	}

	t.Run("customerID_is_present", func(t *testing.T) {
		customerID := responseDTO.CustomerID

		if customerID == "" {
			t.Errorf("Expected customerID to be present")
		}
	})

	t.Run("accountID_is_present", func(t *testing.T) {
		accountID := responseDTO.AccountID

		if accountID == "" {
			t.Errorf("Expected accountID to be present")
		}
	})

	t.Run("last_modified_time_after_created_time", func(t *testing.T) {
		createdTime, _ := time.Parse(time.RFC3339, responseDTO.CreatedTime)
		lastModifiedTime, _ := time.Parse(time.RFC3339, responseDTO.LastModifiedTime)

		if lastModifiedTime.Before(createdTime) {
			t.Errorf("lastModifiedTime %s cannot be  less before createdTime %s", createdTime, lastModifiedTime)
		}
	})
}
