package record

type Impression struct {
	Identifier int64  `json:"identifier"`
	Content    string `json:"content"`
	Source     string `json:"source"`
	CreatedAt  string `json:"created_at"`
}
