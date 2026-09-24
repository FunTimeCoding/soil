package mutation

type Mutant struct {
	Type        string `json:"type"`
	Status      string `json:"status"`
	Line        int    `json:"line"`
	Mutator     string `json:"mutator"`
	Original    string `json:"original_code"`
	Replacement string `json:"replacement_code"`
}
