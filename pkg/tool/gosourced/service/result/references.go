package result

type References struct {
	Symbol    string      `json:"symbol"`
	Total     int         `json:"total"`
	More      int         `json:"more"`
	Locations []*Location `json:"locations"`
}
