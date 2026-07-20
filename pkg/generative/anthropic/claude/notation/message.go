package notation

import "encoding/json"

type Message struct {
	Identifier string          `json:"id"`
	Role       string          `json:"role"`
	Content    json.RawMessage `json:"content"`
	Model      string          `json:"model"`
	Usage      *MessageUsage   `json:"usage"`
}
