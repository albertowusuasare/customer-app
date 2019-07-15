package app

import (
	"github.com/albertowusuasare/customer-app/internal/workflow"
)

// CustomerAppFunc creates a new instance of the customer application
type CustomerAppFunc func() Customer

// Customer encapsulates all the workflows for the application
type Customer struct {
	CreateWf         workflow.CreateFunc
	RetrieveSingleWf workflow.RetrieveSingleFunc
	RetrieveMultiWf  workflow.RetrieveMultiFunc
	UpdateWf         workflow.UpdateFunc
	RemoveWf         workflow.RemoveFunc
}
