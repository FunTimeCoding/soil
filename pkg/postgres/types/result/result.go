package result

type Result struct {
	Columns []string `json:"columns"`
	Rows    [][]any  `json:"rows"`
}
