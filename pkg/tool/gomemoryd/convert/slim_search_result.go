package convert

type SlimSearchResult struct {
	Identifier       int64    `json:"identifier"`
	Name             string   `json:"name"`
	Content          string   `json:"content"`
	Description      string   `json:"description,omitempty"`
	Scope            string   `json:"scope,omitempty"`
	Tags             []string `json:"tags,omitempty"`
	Rank             float64  `json:"rank"`
	ParentIdentifier *int64   `json:"parent_identifier,omitempty"`
	ParentName       string   `json:"parent_name,omitempty"`
}
