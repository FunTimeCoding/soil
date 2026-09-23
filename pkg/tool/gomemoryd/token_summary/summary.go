package token_summary

type Summary struct {
	Statistic   []*Statistic `json:"statistic"`
	Block       *Spread      `json:"block"`
	Description *Spread      `json:"description"`
	Withheld    int          `json:"withheld"`
}
