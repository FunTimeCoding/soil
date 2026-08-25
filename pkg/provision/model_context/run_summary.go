package model_context

type RunSummary struct {
	Identifier          uint   `json:"id"`
	CreatedAt           string `json:"created_at"`
	Scope               string `json:"scope,omitempty"`
	TriggerSource       string `json:"trigger_source"`
	DurationMillisecond int64  `json:"duration_ms"`
	Status              string `json:"status"`
	GitHead             string `json:"git_head"`
}
