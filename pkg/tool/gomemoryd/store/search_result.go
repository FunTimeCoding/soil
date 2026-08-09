package store

type SearchResult struct {
	Identifier       int64    `json:"identifier"`
	Name             string   `json:"name"`
	Content          string   `json:"content"`
	Description      string   `json:"description"`
	Type             string   `json:"type"`
	Scope            string   `json:"scope,omitempty"`
	Tags             []string `json:"tags,omitempty"`
	UpdatedAt        string   `json:"updated_at"`
	Rank             float64  `json:"rank"`
	ParentIdentifier *int64   `json:"parent_identifier,omitempty"`
	ParentName       string   `json:"parent_name,omitempty"`
}
