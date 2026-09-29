package argument

type AddNote struct {
	LinkIdentifier int    `json:"link_identifier"`
	Text           string `json:"text"`
}
