package argument

type SearchConversations struct {
	Query string   `json:"query"`
	Kinds []string `json:"kinds"`
	Limit float64  `json:"limit"`
}
