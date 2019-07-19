package integration

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/albertowusuasare/customer-app/internal/api"
	"github.com/albertowusuasare/customer-app/internal/app"
	"github.com/albertowusuasare/customer-app/internal/uuid"
)

func TestCreateAPI(t *testing.T) {
	app := app.Inmem()
	ts := httptest.NewServer(api.Handler(app))
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
	requestDTO := api.CreateRequestDTO{}
	UnMarshal(request, &requestDTO)
	responseDTO := api.CreateResponseDTO{}
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
			t.Fatalf("Customer ID not present. got %s", customerID)
		}
	})

	t.Run("CustomerIDValidUUID", func(t *testing.T) {
		customerID := responseDTO.CustomerID
		if !uuid.IsValidUUID(customerID) {
			t.Errorf("CustomerID is not a valid v4 UUID. got %s", customerID)
		}
	})
}

// ValidationErrTestCase is a tabular description of validation error tests
type ValidationErrTestCase struct {
	fieldName,
	scenario,
	requestFile,
	expectedResponseFile string
}

func TestCreateErrorResponse(t *testing.T) {
	app := app.Inmem()
	ts := httptest.NewServer(api.Handler(app))
	defer ts.Close()

	testCases := validationErrorTestCases()
	testDataRelativePath := "../data/validation/create"
	for _, tc := range testCases {
		testName := fmt.Sprintf("%s-%s", tc.fieldName, tc.scenario)
		t.Run(testName, func(t *testing.T) {

			// invoke API
			requestBody, _ := ioutil.ReadFile(fmt.Sprintf("%s/%s", testDataRelativePath, tc.requestFile))
			buffer := bytes.NewBuffer(requestBody)
			res, err := http.Post(ts.URL+"/customers/", "application/json", buffer)
			if err != nil {
				log.Fatal(err)
			}

			// assert status code
			statusCode := res.StatusCode
			expectedStatusCode := http.StatusNotAcceptable
			if statusCode != http.StatusNotAcceptable {
				t.Fatalf("Status was %d expecting %d", statusCode, expectedStatusCode)
			}

			// assert error body response
			responseBody, _ := ioutil.ReadAll(res.Body)
			errorBody := api.Error{}
			UnMarshal(responseBody, &errorBody)

			expectedErrorBody := api.Error{}
			expectedErrBodyFile, _ := ioutil.ReadFile(fmt.Sprintf("%s/%s", testDataRelativePath, tc.expectedResponseFile))
			UnMarshal(expectedErrBodyFile, &expectedErrorBody)

			if !reflect.DeepEqual(expectedErrorBody, errorBody) {
				t.Errorf("Expected error response payload %+v got %+v", expectedErrorBody, errorBody)
			}

		})
	}
}

func validationErrorTestCases() []ValidationErrTestCase {
	return []ValidationErrTestCase{
		{
			"firstName",
			"empty",
			"create-empty-firstname-request.json",
			"create-empty-firstname-response.json",
		},
		{
			"firstName",
			"alphanumeric",
			"create-nonalphanumeric-firstname-request.json",
			"create-nonalphanumeric-firstname-response.json",
		},
		{
			"firstName",
			"long-length",
			"create-long-length-firstname-request.json",
			"create-long-length-firstname-response.json",
		},
		{
			"lastName",
			"empty",
			"create-empty-lastname-request.json",
			"create-empty-lastname-response.json",
		},
		{
			"lastName",
			"alphanumeric",
			"create-nonalphanumeric-lastname-request.json",
			"create-nonalphanumeric-lastname-response.json",
		},
		{
			"lastName",
			"long-length",
			"create-long-length-lastname-request.json",
			"create-long-length-lastname-response.json",
		},
	}
}
