package model_context

type groupRelation struct {
	SourceIdentifier int64  `json:"source_id"`
	Source           string `json:"source"`
	SourceScope      string `json:"source_scope,omitempty"`
	Type             string `json:"type,omitempty"`
	TargetIdentifier int64  `json:"target_id"`
	Target           string `json:"target"`
	TargetScope      string `json:"target_scope,omitempty"`
}
