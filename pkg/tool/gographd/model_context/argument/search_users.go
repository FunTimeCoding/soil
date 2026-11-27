package argument

type SearchUsers struct {
	Query string `json:"query"`
	Limit int    `json:"limit"`
}
