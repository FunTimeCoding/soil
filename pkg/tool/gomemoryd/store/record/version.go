package record

type Version struct {
	Identifier       int64  `json:"identifier"`
	MemoryIdentifier int64  `json:"memory_identifier"`
	Name             string `json:"name"`
	Content          string `json:"content"`
	Description      string `json:"description"`
	ChangedAt        string `json:"changed_at"`
	ChangeType       string `json:"change_type"`
	Source           string `json:"source"`
}
