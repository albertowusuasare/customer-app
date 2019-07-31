package app

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/albertowusuasare/customer-app/internal/adding"
	queue "github.com/albertowusuasare/customer-app/internal/msg/inmem"
	"github.com/albertowusuasare/customer-app/internal/storage/google"
	"github.com/albertowusuasare/customer-app/internal/storage/inmem"
	"github.com/albertowusuasare/customer-app/internal/uuid"
	"github.com/albertowusuasare/customer-app/internal/workflow"
)

// GoogleApp creates a customer app based on in memory data store
func GoogleApp(ctx context.Context, firestoreClient *firestore.Client) Customer {
	firestoreInsert := google.CreateCustomerDoc(ctx, firestoreClient)
	createWf := workflow.Create(adding.ValidateRequest, uuid.GenV4, firestoreInsert, queue.CustomerAddedPublisher())

	firestoreRetrieve := google.RetrieveCustomerDoc(ctx, firestoreClient)
	retrieveSingleWf := workflow.RetrieveOne(firestoreRetrieve)

	firestoreRetrieveMulti := google.RetrieveCustomerDocs(ctx, firestoreClient)
	retrieveMultiWf := workflow.RetrieveMulti(firestoreRetrieveMulti)
	updateWf := workflow.Update(inmem.UpdateCustomer(), queue.CustomerUpdatedPublisher())
	removeWf := workflow.Remove(inmem.RemoveCustomer(), queue.CustomerRemovedPublisher())

	return Customer{
		CreateWf:         createWf,
		RetrieveSingleWf: retrieveSingleWf,
		RetrieveMultiWf:  retrieveMultiWf,
		UpdateWf:         updateWf,
		RemoveWf:         removeWf,
	}
}
