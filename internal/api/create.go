package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/albertowusuasare/customer-app/internal/adding"
	"github.com/albertowusuasare/customer-app/internal/validation"
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
	CustomerID  string `json:"customerId"`
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
		peristedCustomer, err := wf(createRequest)
		if err != nil {
			handleWorkflowError(err, w)
			return
		}
		response := createResponseDTOFromCustomer(peristedCustomer)
		encodeErr := json.NewEncoder(w).Encode(response)
		if encodeErr != nil {
			log.Fatal(encodeErr)
		}
	}
}

func createRequestFromCreateRequestDTO(dto CreateRequestDTO) adding.UnvalidatedRequest {
	return adding.UnvalidatedRequest{
		FirstName:   dto.FirstName,
		LastName:    dto.LastName,
		NationalID:  dto.NationalID,
		PhoneNumber: dto.PhoneNumber,
		AccountID:   dto.AccountID,
	}
}

func createResponseDTOFromCustomer(peristedCustomer adding.Customer) CreateResponseDTO {
	return CreateResponseDTO{
		CustomerID:  peristedCustomer.CustomerID,
		FirstName:   peristedCustomer.FirstName,
		LastName:    peristedCustomer.LastName,
		NationalID:  peristedCustomer.NationalID,
		PhoneNumber: peristedCustomer.PhoneNumber,
		AccountID:   peristedCustomer.AccountID,
	}
}

func handleWorkflowError(err error, w http.ResponseWriter) {
	if !validation.IsFieldValidationError(err) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fields, _ := validation.GetFailedValidationFields(err)
	params := map[string]string{}

	for k, v := range fields {
		params[validation.RetrieveFieldName(k)] = validation.RetrieveMessage(v)
	}

	errorBody := Error{
		Code:    InvalidRequestBody,
		Message: InputValidationErrorMessage,
		Params:  params,
	}

	w.WriteHeader(http.StatusNotAcceptable)
	b, marshalErr := json.Marshal(errorBody)
	if marshalErr != nil {
		log.Fatal(marshalErr)
	}
	_, wErr := w.Write(b)
	if wErr != nil {
		log.Fatal(wErr)
	}
}
