package wasp

import "encoding/json"

type Message struct {
	Name string
}

func (m *Message) Bytes() ([]byte, error) {
	return json.Marshal(m)
}
