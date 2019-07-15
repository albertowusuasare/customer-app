package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/albertowusuasare/customer-app/internal/retrieving"
	"github.com/albertowusuasare/customer-app/internal/workflow"
)

// CustomerRetrieveResponseDTO represents the json structure for a customer retrieve response
type CustomerRetrieveResponseDTO struct {
	CustomerID       string `json:"customerId"`
	FirstName        string `json:"firstName"`
	LastName         string `json:"lastName"`
	NationalID       string `json:"nationalId"`
	PhoneNumber      string `json:"phoneNumber"`
	AccountID        string `json:"accountId"`
	LastModifiedTime string `json:"lastModifiedTime"`
	CreatedTime      string `json:"createdTime"`
	Version          int    `json:"version"`
}

// CustomerRetrieveErrorDTO is the error body returned to the API caller when an error occurs
type CustomerRetrieveErrorDTO struct {
	Message string `json:"message"`
}

// HandleRetrieveOne returns an http handler for a customer retrieve API call
func HandleRetrieveOne(wf workflow.RetrieveSingleFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		customerID := RetrieveCustomerId(r)
		customer, err := wf(customerID)
		if err != nil {
			handleRetrieveOneWorkflowError(err, w)
			return
		}
		response := customerRetrieveResponseDTOFromCustomer(*customer)
		encodeErr := json.NewEncoder(w).Encode(response)
		if encodeErr != nil {
			log.Fatal(encodeErr)
		}
	}
}

// HandleRetrieveMulti returns an http handler for a customer retrieve multi API call
func HandleRetrieveMulti(wf workflow.RetrieveMultiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		request := retrieving.MultiRequest{}
		log.Printf("Retrieving customers for request=%+v", request)
		customers := wf(request)
		response := customerRetrieveResponseDTOsFromCustomers(customers)
		encodeErr := json.NewEncoder(w).Encode(response)
		if encodeErr != nil {
			log.Fatal(encodeErr)
		}
	}
}

func handleRetrieveOneWorkflowError(err error, w http.ResponseWriter) {
	switch err.(type) {
	case retrieving.CustomerNonExistent:
		{
			w.WriteHeader(http.StatusNotFound)
			errorDTO := CustomerRetrieveErrorDTO{err.Error()}
			b, marshalErr := json.Marshal(errorDTO)
			if marshalErr != nil {
				log.Fatal(marshalErr)
			}
			_, wErr := w.Write(b)
			if wErr != nil {
				log.Fatal(wErr)
			}
		}
	default:
		log.Fatal(err)
	}
}

func customerRetrieveResponseDTOsFromCustomers(customers []retrieving.Customer) []CustomerRetrieveResponseDTO {
	response := make([]CustomerRetrieveResponseDTO, len(customers))
	for i, c := range customers {
		response[i] = customerRetrieveResponseDTOFromCustomer(c)
	}
	return response
}

func customerRetrieveResponseDTOFromCustomer(customer retrieving.Customer) CustomerRetrieveResponseDTO {
	return CustomerRetrieveResponseDTO{
		CustomerID:       customer.CustomerId,
		FirstName:        customer.FirstName,
		LastName:         customer.LastName,
		NationalID:       customer.NationalId,
		PhoneNumber:      customer.PhoneNumber,
		AccountID:        customer.AccountId,
		LastModifiedTime: customer.LastModifiedTime,
		CreatedTime:      customer.CreatedTime,
		Version:          customer.Version,
	}
}
