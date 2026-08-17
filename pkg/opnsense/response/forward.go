package response

type Forward struct {
	Identifier  string `json:"uuid"`
	Enabled     Flag   `json:"enabled"`
	Type        string `json:"type"`
	Domain      string `json:"domain"`
	Server      string `json:"server"`
	Port        string `json:"port"`
	Description string `json:"description"`
}
