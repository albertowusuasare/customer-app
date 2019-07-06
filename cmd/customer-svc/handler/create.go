package handler

import (
	"customer-app/internal/adding"
	"customer-app/internal/workflow"
	"encoding/json"
	"net/http"
)

type CreateRequestDTO struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	NationalId  string `json:"nationalId"`
	PhoneNumber string `json:"phoneNumber"`
	AccountId   string `json:"accountId"`
}

type CreateResponseDTO struct {
	CustomerId  string `json:"customerId"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	NationalId  string `json:"nationalId"`
	PhoneNumber string `json:"phoneNumber"`
	AccountId   string `json:"accountId"`
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
		NationalId:  dto.NationalId,
		PhoneNumber: dto.NationalId,
		AccountId:   dto.AccountId,
	}
}

func createResponseDTOFromPersistedCustomer(peristedCustomer adding.PersistedCustomer) CreateResponseDTO {
	return CreateResponseDTO{
		CustomerId:  peristedCustomer.CustomerId,
		FirstName:   peristedCustomer.FirstName,
		LastName:    peristedCustomer.LastName,
		NationalId:  peristedCustomer.NationalId,
		PhoneNumber: peristedCustomer.PhoneNumber,
		AccountId:   peristedCustomer.AccountId,
	}
}
