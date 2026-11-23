package channel

type Channel struct {
	Identifier  string `json:"identifier"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Membership  string `json:"membership,omitempty"`
	Link        string `json:"link,omitempty"`
}
