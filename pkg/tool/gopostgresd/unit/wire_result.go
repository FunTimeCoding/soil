package unit

type wireResult struct {
	Columns []string `json:"columns"`
	Rows    [][]any  `json:"rows"`
}
