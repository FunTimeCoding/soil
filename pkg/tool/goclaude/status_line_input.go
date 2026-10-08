package goclaude

type StatusLineInput struct {
	SessionIdentifier string                `json:"session_id"`
	Model             StatusLineModel       `json:"model"`
	ContextWindow     StatusLineContext     `json:"context_window"`
	RateLimits        *StatusLineRateLimits `json:"rate_limits"`
}
