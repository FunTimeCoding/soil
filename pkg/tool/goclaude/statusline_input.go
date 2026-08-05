package goclaude

type statuslineInput struct {
	SessionID     string                `json:"session_id"`
	Model         statuslineModel       `json:"model"`
	ContextWindow statuslineContext     `json:"context_window"`
	RateLimits    *statuslineRateLimits `json:"rate_limits"`
}
