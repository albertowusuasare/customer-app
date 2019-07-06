package inmem

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Envelope struct {
	Header Header
	Body   []byte
}

type Header struct {
	Id        string
	TimeStamp string
}

func CreateEnvelope(payload interface{}) Envelope {
	header := Header{
		Id:        uuid.New().String(),
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
