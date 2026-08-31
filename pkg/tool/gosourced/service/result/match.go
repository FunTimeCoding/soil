package result

type Match struct {
	Symbol    string   `json:"symbol"`
	Pattern   string   `json:"pattern"`
	Total     int      `json:"total"`
	Matched   int      `json:"matched"`
	Unmatched []*Group `json:"unmatched"`
}
