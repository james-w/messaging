package shared

import (
	"encoding/json"
)

type Message struct {
	Username string `json:"username"`
	Text     string `json:"text"`
}

func Encode(msg Message) ([]byte, error) {
	return json.Marshal(msg)
}

func Decode(data []byte) (Message, error) {
	var m Message
	err := json.Unmarshal(data, &m)
	return m, err
}
