package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/albertowusuasare/customer-app/internal/adding"
	"github.com/albertowusuasare/customer-app/internal/workflow"
)

// CreateRequestDTO represents the json structure for a customer create request
type CreateRequestDTO struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	NationalID  string `json:"nationalId"`
	PhoneNumber string `json:"phoneNumber"`
	AccountID   string `json:"accountId"`
}

// CreateResponseDTO represents the json structure for a customer create response
type CreateResponseDTO struct {
	CustomerID  string `json:"customerId"` //rename this ID?
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	NationalID  string `json:"nationalId"`
	PhoneNumber string `json:"phoneNumber"`
	AccountID   string `json:"accountId"`
}

// HandleCreate returns an http handler for the customer create API call
func HandleCreate(wf workflow.CreateFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var requestDTO CreateRequestDTO
		err := decoder.Decode(&requestDTO)
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "application/json")
		createRequest := createRequestFromCreateRequestDTO(requestDTO)
		peristedCustomer := wf(createRequest)
		response := createResponseDTOFromPersistedCustomer(peristedCustomer)
		encodeErr := json.NewEncoder(w).Encode(response)
		if encodeErr != nil {
			log.Fatal(encodeErr)
		}
	}
}

func createRequestFromCreateRequestDTO(dto CreateRequestDTO) adding.Request {
	return adding.Request{
		FirstName:   dto.FirstName,
		LastName:    dto.LastName,
		NationalId:  dto.NationalID,
		PhoneNumber: dto.PhoneNumber,
		AccountId:   dto.AccountID,
	}
}

func createResponseDTOFromPersistedCustomer(peristedCustomer adding.PersistedCustomer) CreateResponseDTO {
	return CreateResponseDTO{
		CustomerID:  peristedCustomer.CustomerId,
		FirstName:   peristedCustomer.FirstName,
		LastName:    peristedCustomer.LastName,
		NationalID:  peristedCustomer.NationalId,
		PhoneNumber: peristedCustomer.PhoneNumber,
		AccountID:   peristedCustomer.AccountId,
	}
}
