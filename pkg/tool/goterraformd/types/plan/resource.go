package plan

type Resource struct {
	Address     string         `json:"address"`
	Type        string         `json:"type"`
	Expressions map[string]any `json:"expressions"`
}
