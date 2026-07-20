package result

type Location struct {
	File    string `json:"file"`
	Line    int    `json:"line"`
	Package string `json:"package"`
}
