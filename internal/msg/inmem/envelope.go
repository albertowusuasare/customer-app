package inmem

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// An Envelope is a representation of  the data sent to a message broker for an event
type Envelope struct {
	Header Header
	Body   []byte
}

// A Header is a representation of an event meta data.
type Header struct {
	ID        string
	TimeStamp string
}

// CreateEnvelope creates an evelope for any data payload
func CreateEnvelope(payload interface{}) Envelope {
	header := Header{
		ID:        uuid.New().String(),
		TimeStamp: time.Now().Format(time.RFC3339),
	}
	body := marshal(payload)

	return Envelope{
		Header: header,
		Body:   body,
	}
}

func marshal(event interface{}) []byte {
	b, err := json.Marshal(event)
	if err != nil {
		fmt.Println(err)
	}
	return b
}
