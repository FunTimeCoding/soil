package store

type Memory struct {
	Identifier       int64             `json:"identifier"`
	Name             string            `json:"name"`
	Content          string            `json:"content"`
	Description      string            `json:"description"`
	Type             string            `json:"type"`
	Scope            string            `json:"scope,omitempty"`
	CreatedAt        string            `json:"created_at"`
	UpdatedAt        string            `json:"updated_at"`
	IsActive         bool              `json:"is_active"`
	Tags             []string          `json:"tags,omitempty"`
	Metadata         map[string]string `json:"metadata,omitempty"`
	ParentIdentifier *int64            `json:"parent_identifier,omitempty"`
	ProvenanceFile   string            `json:"provenance_file,omitempty"`
	ProvenanceAnchor string            `json:"provenance_anchor,omitempty"`
	ProvenanceHash   string            `json:"provenance_hash,omitempty"`
	Ordinal          int               `json:"ordinal,omitempty"`
}
