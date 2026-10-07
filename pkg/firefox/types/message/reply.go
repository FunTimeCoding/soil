package message

import "encoding/json"

type Reply struct {
	Result     json.RawMessage `json:"result,omitempty"`
	Error      string          `json:"error,omitempty"`
	Identifier int             `json:"id"`
}
