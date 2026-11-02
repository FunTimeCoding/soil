package transcript

import "encoding/json"

type Message struct {
	Content json.RawMessage `json:"content"`
}
