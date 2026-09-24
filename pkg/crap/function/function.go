package function

type Function struct {
	File       string `json:"file"`
	Line       int    `json:"line"`
	EndLine    int    `json:"end_line"`
	Package    string `json:"package"`
	Receiver   string `json:"receiver"`
	Name       string `json:"name"`
	Complexity int    `json:"complexity"`
}
