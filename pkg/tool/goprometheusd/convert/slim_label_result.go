package convert

type SlimLabelResult struct {
	Values   []string `json:"values"`
	Warnings []string `json:"warnings,omitempty"`
	Notices  []string `json:"notices,omitempty"`
}
