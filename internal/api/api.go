package api

import (
	"net/http"

	"github.com/albertowusuasare/customer-app/internal/workflow"
)

// CreateHandlerFunc returns an http handler for a customer create API call
type CreateHandlerFunc func(wf workflow.CreateFunc) http.HandlerFunc

// RetriveOneHandlerFunc returns an http handler for a customer retrieve API call
type RetriveOneHandlerFunc func(wf workflow.RetrieveSingleFunc) http.HandlerFunc

// RetriveMultiHandlerFunc returns an http handler for a customer retrieve multi API call
type RetriveMultiHandlerFunc func(wf workflow.RetrieveMultiFunc) http.HandlerFunc

// UpdateHandlerFunc returns an http handler for a customer update API call
type UpdateHandlerFunc func(wf workflow.UpdateFunc) http.HandlerFunc

// RemoveHandlerFunc returns an http handler for a customer remove API call
type RemoveHandlerFunc func(wf workflow.RemoveFunc) http.HandlerFunc
