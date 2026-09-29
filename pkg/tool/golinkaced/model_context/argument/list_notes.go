package argument

type ListNotes struct {
	LinkIdentifier int `json:"link_identifier"`
	Limit          int `json:"limit"`
	Offset         int `json:"offset"`
}
