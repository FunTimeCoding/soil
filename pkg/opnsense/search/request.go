package search

type Request struct {
	Current      int    `json:"current"`
	RowCount     int    `json:"rowCount"`
	SearchPhrase string `json:"searchPhrase"`
}
