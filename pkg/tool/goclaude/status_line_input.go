package goclaude

type statusLineInput struct {
	SessionID     string                `json:"session_id"`
	Model         statusLineModel       `json:"model"`
	ContextWindow statusLineContext     `json:"context_window"`
	RateLimits    *statusLineRateLimits `json:"rate_limits"`
}
