package app

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/albertowusuasare/customer-app/internal/adding"
	"github.com/albertowusuasare/customer-app/internal/validation"
	"github.com/albertowusuasare/customer-app/internal/workflow"
	"github.com/pkg/errors"
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

// CreateHandlerFunc returns an http handler for a customer create API call
type CreateHandlerFunc func(wf workflow.CreateFunc) http.HandlerFunc

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
		customer, err := wf(createRequest)
		if err != nil {
			handleWorkflowError(err, w)
			return
		}
		response := createResponseDTOFromCustomer(*customer)
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

func createResponseDTOFromCustomer(c adding.Customer) CreateResponseDTO {
	return CreateResponseDTO{
		CustomerID:  string(c.RetrieveCustomerID()),
		FirstName:   string(c.RetrieveFirstName()),
		LastName:    string(c.RetrieveLastName()),
		NationalID:  string(c.RetrieveNationalID()),
		PhoneNumber: string(c.RetrievePhoneNumber()),
		AccountID:   string(c.RetrieveAccountID()),
	}
}

func handleWorkflowError(err error, w http.ResponseWriter) {
	cause := errors.Cause(err)
	if !validation.IsFieldValidationError(cause) {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fields, _ := validation.GetFailedValidationFields(cause)
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
