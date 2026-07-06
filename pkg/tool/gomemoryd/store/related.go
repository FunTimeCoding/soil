package store

type Related struct {
	Identifier  int64    `json:"identifier"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Tags        []string `json:"tags,omitempty"`
}
