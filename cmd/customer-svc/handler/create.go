package handler

import (
	"encoding/json"
	"net/http"

	"github.com/albertowusuasare/customer-app/internal/adding"
	"github.com/albertowusuasare/customer-app/internal/workflow"
)

type CreateRequestDTO struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	NationalID  string `json:"nationalId"`
	PhoneNumber string `json:"phoneNumber"`
	AccountID   string `json:"accountId"`
}

type CreateResponseDTO struct {
	CustomerID  string `json:"customerId"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	NationalID  string `json:"nationalId"`
	PhoneNumber string `json:"phoneNumber"`
	AccountID   string `json:"accountId"`
}

type CreateHandler struct {
	Workflow workflow.CreateFunc
}

func (handler CreateHandler) Handle(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var requestDTO CreateRequestDTO
	err := decoder.Decode(&requestDTO)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	createRequest := createRequestFromCreateRequestDTO(requestDTO)
	peristedCustomer := handler.Workflow(createRequest)
	response := createResponseDTOFromPersistedCustomer(peristedCustomer)
	json.NewEncoder(w).Encode(response)
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
