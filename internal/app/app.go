package app

import (
	"github.com/albertowusuasare/customer-app/internal/workflow"
)

// Customer encapsulates all the workflows for the application
type Customer struct {
	CreateWf         workflow.CreateFunc
	RetrieveSingleWf workflow.RetrieveSingleFunc
	RetrieveMultiWf  workflow.RetrieveMultiFunc
	UpdateWf         workflow.UpdateFunc
	RemoveWf         workflow.RemoveFunc
}
