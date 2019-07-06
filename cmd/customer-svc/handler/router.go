package handler

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

func Router(
	customerCreateHandler CreateHandler,
	retrieveOneHandler RetrieveOneHandler,
	retrieveMultiHandler RetrieveMultiHandler,
	updateHandler UpdateHandler,
	removeHandler RemoveHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		method := r.Method

		if method == "POST" {
			customerCreateHandler.Handle(w, r)
			return
		}

		if method == "GET" {
			handleGetCall(w, r, retrieveOneHandler, retrieveMultiHandler)
			return
		}

		if method == "PUT" {
			updateHandler.Handle(w, r)
			return
		}

		if method == "DELETE" {
			removeHandler.Handle(w, r)
			return
		}

		log.Panic(fmt.Sprintf("httpMethod=%s not supported\n", method))
	}
}

func handleGetCall(
	w http.ResponseWriter,
	r *http.Request,
	retrieveOneHandler RetrieveOneHandler,
	retrieveMultiHandler RetrieveMultiHandler) {
	uri := r.RequestURI
	customerIdPresent := hasCustomerId(uri)
	if customerIdPresent {
		retrieveOneHandler.Handle(w, r)
		return
	}
	retrieveMultiHandler.Handle(w, r)
}

func hasCustomerId(requestURI string) bool {
	customerIdValue := strings.Split(requestURI, "/")[2]
	_, err := uuid.Parse(customerIdValue)
	return err == nil
}
