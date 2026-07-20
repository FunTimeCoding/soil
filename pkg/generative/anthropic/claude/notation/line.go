package notation

import "encoding/json"

type Line struct {
	Type          string          `json:"type"`
	Slug          string          `json:"slug"`
	Timestamp     string          `json:"timestamp"`
	Session       string          `json:"sessionId"`
	Request       string          `json:"requestId"`
	WorkDirectory string          `json:"cwd"`
	Branch        string          `json:"gitBranch"`
	Meta          bool            `json:"isMeta"`
	Message       json.RawMessage `json:"message"`
}
