package result

type Apply struct {
	Symbol      string   `json:"symbol"`
	Pattern     string   `json:"pattern"`
	Replacement string   `json:"replacement"`
	Total       int      `json:"total"`
	Matched     int      `json:"matched"`
	Rewritten   int      `json:"rewritten"`
	Applied     bool     `json:"applied"`
	Refusal     string   `json:"refusal,omitempty"`
	Unmatched   []*Group `json:"unmatched"`
	Refused     []*Group `json:"refused"`
}
