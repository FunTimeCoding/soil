package argument

type SearchItems struct {
	Q       string   `json:"q"`
	Types   []string `json:"types"`
	Page    int      `json:"page"`
	PerPage int      `json:"per_page"`
}
