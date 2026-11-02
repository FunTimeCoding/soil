package transcript

import "encoding/json"

type Block struct {
	Type      string          `json:"type"`
	Name      string          `json:"name"`
	Identifier string         `json:"id"`
	ToolUse   string          `json:"tool_use_id"`
	Input     json.RawMessage `json:"input"`
	Content   json.RawMessage `json:"content"`
}
