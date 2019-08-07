package app

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/albertowusuasare/customer-app/internal/updating"
	"github.com/albertowusuasare/customer-app/internal/workflow"
)

// UpdateRequestDTO represents the json structure for a customer update request
type UpdateRequestDTO struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	NationalID  string `json:"nationalId"`
	PhoneNumber string `json:"phoneNumber"`
}

// UpdateResponseDTO represents the json structure for a customer update response
type UpdateResponseDTO struct {
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

// RemoveHandlerFunc returns an http handler for a customer remove API call
type RemoveHandlerFunc func(wf workflow.RemoveFunc) http.HandlerFunc

// UpdateHandlerFunc returns an http handler for a customer update API call
type UpdateHandlerFunc func(wf workflow.UpdateFunc) http.HandlerFunc

// HandleUpdate returns an http handler for a customer update API call
func HandleUpdate(wf workflow.UpdateFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var requestDTO UpdateRequestDTO
		err := decoder.Decode(&requestDTO)
		if err != nil {
			panic(err)
		}
		customerID := RetrieveCustomerID(r)
		w.Header().Set("Content-Type", "application/json")
		request := updateRequestFromUpdateRequestDTO(customerID, requestDTO)
		log.Printf("Updating customer for request=%+v", request)
		updatedCustomer := wf(request)
		response := updateResponseDTOFromUpdatedCustomer(updatedCustomer)
		encodeErr := json.NewEncoder(w).Encode(response)
		if encodeErr != nil {
			log.Fatal(encodeErr)
		}
	}
}

func updateRequestFromUpdateRequestDTO(customerID string, dto UpdateRequestDTO) updating.Request {
	return updating.Request{
		CustomerID:  customerID,
		FirstName:   dto.FirstName,
		LastName:    dto.LastName,
		NationalID:  dto.NationalID,
		PhoneNumber: dto.PhoneNumber,
	}
}

func updateResponseDTOFromUpdatedCustomer(customer updating.UpdatedCustomer) UpdateResponseDTO {
	return UpdateResponseDTO{
		CustomerID:       customer.CustomerID,
		FirstName:        customer.FirstName,
		LastName:         customer.LastName,
		NationalID:       customer.NationalID,
		PhoneNumber:      customer.PhoneNumber,
		AccountID:        customer.AccountID,
		LastModifiedTime: customer.LastModifiedTime,
		CreatedTime:      customer.CreatedTime,
		Version:          customer.Version,
	}
}
