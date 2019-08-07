package google

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/albertowusuasare/customer-app/internal/adding"
	"github.com/albertowusuasare/customer-app/internal/app"
	queue "github.com/albertowusuasare/customer-app/internal/msg/inmem"
	"github.com/albertowusuasare/customer-app/internal/storage/google"
	"github.com/albertowusuasare/customer-app/internal/uuid"
	"github.com/albertowusuasare/customer-app/internal/workflow"
)

// App creates a customer app based on in memory data store
func App(ctx context.Context, firestoreClient *firestore.Client) app.Customer {
	firestoreInsert := google.CreateCustomerDoc(ctx, firestoreClient)
	createWf := workflow.Create(adding.ValidateRequest, uuid.GenV4, firestoreInsert, queue.CustomerAddedPublisher())

	firestoreRetrieve := google.RetrieveCustomerDoc(ctx, firestoreClient)
	retrieveSingleWf := workflow.RetrieveOne(firestoreRetrieve)

	firestoreRetrieveMulti := google.RetrieveCustomerDocs(ctx, firestoreClient)
	retrieveMultiWf := workflow.RetrieveMulti(firestoreRetrieveMulti)

	firestoreUpdate := google.UpdateCustomerDoc(ctx, firestoreClient)
	updateWf := workflow.Update(firestoreUpdate, queue.CustomerUpdatedPublisher())

	firestoreRemove := google.DeleteCustomerDoc(ctx, firestoreClient)
	removeWf := workflow.Remove(firestoreRemove, queue.CustomerRemovedPublisher())

	return app.Customer{
		CreateWf:         createWf,
		RetrieveSingleWf: retrieveSingleWf,
		RetrieveMultiWf:  retrieveMultiWf,
		UpdateWf:         updateWf,
		RemoveWf:         removeWf,
	}
}
