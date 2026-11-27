package outgoing

type Member struct {
	Type    string   `json:"@odata.type"`
	Roles   []string `json:"roles"`
	Binding string   `json:"user@odata.bind"`
}
