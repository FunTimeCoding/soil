package response

type Alias struct {
	Identifier  string `json:"uuid"`
	Enabled     Flag   `json:"enabled"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Content     string `json:"content"`
	Description string `json:"description"`
}
