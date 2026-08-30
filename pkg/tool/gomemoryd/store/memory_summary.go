package store

type MemorySummary struct {
	Identifier       int64    `json:"identifier"`
	Name             string   `json:"name"`
	Description      string   `json:"description"`
	Type             string   `json:"type"`
	Scope            string   `json:"scope,omitempty"`
	Tags             []string `json:"tags,omitempty"`
	UpdatedAt        string   `json:"updated_at"`
	ParentIdentifier *int64   `json:"parent_identifier,omitempty"`
	Ordinal          int      `json:"ordinal,omitzero"`
	Children         []string `json:"children,omitempty"`
}
