package token_summary

type Statistic struct {
	Identifier  int64    `json:"identifier"`
	Name        string   `json:"name"`
	Tags        []string `json:"tags,omitempty"`
	Block       int      `json:"block"`
	Description int      `json:"description"`
	Hidden      bool     `json:"hidden"`
}
