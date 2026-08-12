package goclaude

type statusLineContext struct {
	UsedPercentage float64 `json:"used_percentage"`
	WindowSize     int     `json:"context_window_size"`
}
