package goclaude

type statuslineContext struct {
	UsedPercentage float64 `json:"used_percentage"`
	WindowSize     int     `json:"context_window_size"`
}
