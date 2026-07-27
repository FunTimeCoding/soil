package store

type SourcedMemory struct {
	Identifier       int64  `json:"identifier"`
	Name             string `json:"name"`
	ParentIdentifier *int64 `json:"parent_identifier,omitempty"`
	ProvenanceFile   string `json:"provenance_file"`
	ProvenanceAnchor string `json:"provenance_anchor"`
	ProvenanceHash   string `json:"provenance_hash"`
	Ordinal          int    `json:"ordinal"`
}
