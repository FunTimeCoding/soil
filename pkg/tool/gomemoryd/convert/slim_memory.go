package convert

type SlimMemory struct {
	Identifier       int64             `json:"identifier"`
	Name             string            `json:"name"`
	Content          string            `json:"content,omitempty"`
	Description      string            `json:"description,omitempty"`
	Tags             []string          `json:"tags,omitempty"`
	Metadata         map[string]string `json:"metadata,omitempty"`
	ParentIdentifier *int64            `json:"parent_identifier,omitempty"`
	Ordinal          int               `json:"ordinal,omitempty"`
}
