package goclaude

type StatusLineContext struct {
	UsedPercentage float64 `json:"used_percentage"`
	WindowSize     int     `json:"context_window_size"`
}
