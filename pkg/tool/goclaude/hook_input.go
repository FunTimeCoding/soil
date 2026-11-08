package goclaude

type HookInput struct {
	SessionIdentifier string `json:"session_id"`
	Reason            string `json:"reason"`
	HookEventName     string `json:"hook_event_name"`
	Trigger           string `json:"trigger"`
}
