package plan

type ValueResource struct {
	Address string         `json:"address"`
	Type    string         `json:"type"`
	Values  map[string]any `json:"values"`
}
