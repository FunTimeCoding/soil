package argument

type Update struct {
	Identifier string   `json:"identifier"`
	Fields     []string `json:"fields"`
}
