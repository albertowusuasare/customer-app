package integration

import (
	"encoding/json"
	"io"
	"log"
)

// UnMarshal deserializes data into value
func UnMarshal(data []byte, value interface{}) {
	err := json.Unmarshal(data, &value)
	if err != nil {
		log.Println(err)
	}
}

// DecodeData deserializes the data in r to the value
func DecodeData(r io.Reader, value interface{}) {
	if err := json.NewDecoder(r).Decode(&value); err == io.EOF {
		return
	}
}
