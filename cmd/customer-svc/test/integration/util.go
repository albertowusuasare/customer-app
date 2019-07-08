package integration

import (
	"encoding/json"
	"log"
)

// UnMarshal deserializes data into value
func UnMarshal(data []byte, value interface{}) {
	err := json.Unmarshal(data, &value)
	if err != nil {
		log.Println(err)
	}
}
