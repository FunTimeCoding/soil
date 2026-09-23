package token_summary

type Spread struct {
	Median      int `json:"median"`
	Ninetieth   int `json:"ninetieth"`
	NinetyNinth int `json:"ninety_ninth"`
	Maximum     int `json:"maximum"`
	Total       int `json:"total"`
}
