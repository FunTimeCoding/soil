package store

type Related struct {
	Identifier  int64    `json:"identifier"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Scope       string   `json:"scope,omitempty"`
	Type        string   `json:"type,omitempty"`
	Tags        []string `json:"tags,omitempty"`
}
