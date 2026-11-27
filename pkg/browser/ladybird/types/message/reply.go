package message

import "encoding/json"

type Reply struct {
	Type       string          `json:"type"`
	Identifier int             `json:"id,omitzero"`
	Result     json.RawMessage `json:"result,omitempty"`
	Error      string          `json:"error,omitempty"`
	Message    string          `json:"message,omitempty"`
	Method     string          `json:"method,omitempty"`
	Parameters json.RawMessage `json:"params,omitempty"`
}
