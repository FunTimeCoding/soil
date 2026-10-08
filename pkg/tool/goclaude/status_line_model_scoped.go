package goclaude

type StatusLineModelScoped struct {
	DisplayName string   `json:"display_name"`
	Utilization *float64 `json:"utilization"`
	ResetsAt    *string  `json:"resets_at"`
}
