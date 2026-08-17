package response

type Blocklist struct {
	Identifier  string `json:"uuid"`
	Enabled     Flag   `json:"enabled"`
	Type        string `json:"type"`
	Description string `json:"description"`
}
