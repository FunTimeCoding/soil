package basic

type commandRequest struct {
	Client     string   `json:"client"`
	Target     string   `json:"tgt,omitempty"`
	Function   string   `json:"fun"`
	Arguments  []string `json:"arg,omitempty"`
	Match      string   `json:"match,omitempty"`
	TargetType string   `json:"tgt_type,omitempty"`
	FullReturn bool     `json:"full_return,omitzero"`
}
