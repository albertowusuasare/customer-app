package google

import (
	"context"
	"net/http"

	"github.com/albertowusuasare/customer-app/internal/msg"

	"cloud.google.com/go/firestore"
	"cloud.google.com/go/pubsub"
	"github.com/albertowusuasare/customer-app/app"
	"github.com/albertowusuasare/customer-app/internal/adding"
	pub "github.com/albertowusuasare/customer-app/internal/msg/google"
	"github.com/albertowusuasare/customer-app/internal/storage/google"
	"github.com/albertowusuasare/customer-app/internal/uuid"
	"github.com/albertowusuasare/customer-app/internal/workflow"
)

// CreateHTTPHandler returns a google implementation for handling customer create http calls
func CreateHTTPHandler(
	ctx context.Context,
	firestoreClient *firestore.Client,
	pubsubClient *pubsub.Client) http.HandlerFunc {
	firestoreInsert := google.CreateCustomerDoc(ctx, firestoreClient)
	publisher := pub.CustomerAddedPubisher(ctx, pubsubClient)
	createWf := workflow.Create(adding.ValidateRequest, uuid.GenV4, firestoreInsert, msg.LoggingCustomerAddedPublisher(publisher))
	return app.HandleCreate(createWf)
}

// RemoveHTTPHandler returns a google implementation for handling customer remove http calls
func RemoveHTTPHandler(
	ctx context.Context,
	firestoreClient *firestore.Client,
	pubsubClient *pubsub.Client) http.HandlerFunc {
	firestoreRemove := google.DeleteCustomerDoc(ctx, firestoreClient)
	publisher := pub.CustomerRemovedPubisher(ctx, pubsubClient)
	removeWf := workflow.Remove(firestoreRemove, msg.LoggingCustomerRemovedPublisher(publisher))
	return app.HandleRemove(removeWf)
}

// RetrieveOneHTTPHandler returns a google implementation for handling single customer retrieval http calls
func RetrieveOneHTTPHandler(ctx context.Context, firestoreClient *firestore.Client) http.HandlerFunc {
	firestoreRetrieve := google.RetrieveCustomerDoc(ctx, firestoreClient)
	retrieveSingleWf := workflow.RetrieveOne(firestoreRetrieve)
	return app.HandleRetrieveOne(retrieveSingleWf)
}

// RetrieveMultiHTTPHander returns a google implementation for handling multi customer retrieval http calls
func RetrieveMultiHTTPHander(ctx context.Context, firestoreClient *firestore.Client) http.HandlerFunc {
	firestoreRetrieveMulti := google.RetrieveCustomerDocs(ctx, firestoreClient)
	retrieveMultiWf := workflow.RetrieveMulti(firestoreRetrieveMulti)
	return app.HandleRetrieveMulti(retrieveMultiWf)
}

// UpdateHTTPHandler returns a google implementation for handling customer update http calls
func UpdateHTTPHandler(
	ctx context.Context,
	firestoreClient *firestore.Client,
	pubsubClient *pubsub.Client) http.HandlerFunc {
	firestoreUpdate := google.UpdateCustomerDoc(ctx, firestoreClient)
	publisher := pub.CustomerUpdatedPubisher(ctx, pubsubClient)
	updateWf := workflow.Update(firestoreUpdate, msg.LoggingCustomerUpdatedPublisher(publisher))
	return app.HandleUpdate(updateWf)
}

// App creates a customer app based on in memory data store
func App(ctx context.Context, firestoreClient *firestore.Client, pubsubClient *pubsub.Client) app.StandAlone {

	createHandler := CreateHTTPHandler(ctx, firestoreClient, pubsubClient)
	retrieveOneHandler := RetrieveOneHTTPHandler(ctx, firestoreClient)
	retrieveMultiHandler := RetrieveMultiHTTPHander(ctx, firestoreClient)
	updateHandler := UpdateHTTPHandler(ctx, firestoreClient, pubsubClient)
	removeHandler := RemoveHTTPHandler(ctx, firestoreClient, pubsubClient)

	return app.StandAlone{
		CreateHandler:        createHandler,
		RetrieveOneHandler:   retrieveOneHandler,
		RetrieveMultiHandler: retrieveMultiHandler,
		UpdateHandler:        updateHandler,
		RemoveHandler:        removeHandler,
	}
}
