package goclaude

type HookInput struct {
	SessionIdentifier string `json:"session_id"`
	Reason            string `json:"reason"`
}
