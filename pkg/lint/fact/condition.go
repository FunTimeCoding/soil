package fact

type Condition struct {
	Kind     string `json:"kind"`
	Callee   string `json:"callee"`
	Argument int    `json:"argument"`
	Holder   string `json:"holder"`
	Field    string `json:"field"`
}
