package app

import (
	"customer-app/internal/workflow"
)

type Customer struct {
	CreateWf         workflow.CreateFunc
	RetrieveSingleWf workflow.RetrieveSingleFunc
	RetrieveMultiWf  workflow.RetrieveMultiFunc
	UpdateWf         workflow.UpdateFunc
	RemoveWf         workflow.RemoveFunc
}
