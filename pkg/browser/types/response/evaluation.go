package response

import "encoding/json"

type Evaluation struct {
	Type      string          `json:"type"`
	Result    json.RawMessage `json:"result"`
	Exception *Exception      `json:"exceptionDetails"`
}
