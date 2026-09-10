package goclaude

type hookInput struct {
	SessionIdentifier string `json:"session_id"`
	Reason            string `json:"reason"`
}
