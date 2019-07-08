package integration

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/albertowusuasare/customer-app/cmd/customer-svc/handler"
	"github.com/albertowusuasare/customer-app/cmd/customer-svc/pkg"
	"github.com/albertowusuasare/customer-app/internal/uuid"
)

func TestCreateAPI(t *testing.T) {
	app := pkg.InmemApp()
	ts := httptest.NewServer(handler.Handle(app))
	defer ts.Close()

	requestBody, _ := ioutil.ReadFile("../data/create-request.json")
	buffer := bytes.NewBuffer(requestBody)
	res, err := http.Post(ts.URL+"/customers/", "application/json", buffer)
	if err != nil {
		log.Fatal(err)
	}

	if res.Status != "200 OK" {
		t.Fatalf("Status was %s expecting 200", res.Status)
	}

	responseBody, _ := ioutil.ReadAll(res.Body)
	testExpectedResponse(t, requestBody, responseBody)

	defer res.Body.Close()
}

func testExpectedResponse(t *testing.T, request []byte, response []byte) {
	requestDTO := handler.CreateRequestDTO{}
	UnMarshal(request, &requestDTO)
	responseDTO := handler.CreateResponseDTO{}
	UnMarshal(response, &responseDTO)

	var responseFieldsTests = []struct {
		fieldName,
		actual,
		expected string
	}{
		{
			"firstName",
			responseDTO.FirstName,
			requestDTO.FirstName,
		},
		{
			"lastName",
			responseDTO.LastName,
			requestDTO.LastName,
		},
		{
			"nationalId",
			responseDTO.NationalID,
			requestDTO.NationalID,
		},
		{
			"phoneNumber",
			responseDTO.PhoneNumber,
			requestDTO.PhoneNumber,
		},
		{
			"accountId",
			responseDTO.AccountID,
			requestDTO.AccountID,
		},
	}

	for _, tt := range responseFieldsTests {
		t.Run(tt.fieldName, func(t *testing.T) {
			a := tt.actual
			e := tt.expected
			if a != e {
				t.Errorf("got %s want %s", a, e)
			}
		})
	}

	customerID := responseDTO.CustomerID

	t.Run("CustomerIdPresent", func(t *testing.T) {
		if customerID == "" {
			t.Fatalf("Customer id not present. got %s", customerID)
		}
	})

	t.Run("CustomerIdValidUUID", func(t *testing.T) {
		customerID := responseDTO.CustomerID
		if !uuid.IsValidUUID(customerID) {
			t.Errorf("CustomerID is not a valid v4 UUID. got %s", customerID)
		}
	})
}
