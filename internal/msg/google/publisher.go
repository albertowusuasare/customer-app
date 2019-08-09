package google

import (
	"context"
	"log"

	"cloud.google.com/go/pubsub"
	"github.com/albertowusuasare/customer-app/internal/adding"
	"github.com/albertowusuasare/customer-app/internal/msg"
	"github.com/albertowusuasare/customer-app/internal/updating"
	"github.com/pkg/errors"
)

const customerAddedTopic = "customer-add"
const customerUpdatedTopic = "customer-update"
const customerRemovedTopic = "customer-remove"

// CustomerAddedPubisher returns a pubsub implementation of customer added publishing
func CustomerAddedPubisher(ctx context.Context, client *pubsub.Client) msg.CustomerAddedPublisherFunc {
	return func(c *adding.Customer) msg.Response {
		return publishMessage(ctx, client, customerAddedTopic, string(c.RetrieveCustomerID()))
	}
}

func publishMessage(ctx context.Context, client *pubsub.Client, topicName, cID string) msg.Response {
	topic, topicErr := getTopic(ctx, client, topicName)
	if topicErr != nil {
		log.Fatal(topicErr)
	}
	id, err := topic.Publish(ctx, &pubsub.Message{Data: []byte(cID)}).Get(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return msg.Response{MessageID: id, Acknowledged: true}
}

func getTopic(ctx context.Context, client *pubsub.Client, topicName string) (*pubsub.Topic, error) {
	topic := client.Topic(topicName)
	ok, err := topic.Exists(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "Error checking if topic exists")
	}
	if !ok {
		newTopic, createErr := client.CreateTopic(ctx, topicName)
		if createErr != nil {
			return nil, errors.Wrap(err, "Error creating topic")
		}
		return newTopic, nil
	}
	return topic, nil
}

// CustomerUpdatedPubisher returns a pubsub implementation of customer updated publishing
func CustomerUpdatedPubisher(ctx context.Context, client *pubsub.Client) msg.CustomerUpdatedPublisherFunc {
	return func(c updating.UpdatedCustomer) msg.Response {
		return publishMessage(ctx, client, customerUpdatedTopic, c.CustomerID)
	}
}

// CustomerRemovedPubisher returns a pubsub implementation of customer removed publishing
func CustomerRemovedPubisher(ctx context.Context, client *pubsub.Client) msg.CustomerRemovedPublisherFunc {
	return func(cID string) msg.Response {
		return publishMessage(ctx, client, customerRemovedTopic, cID)
	}
}
