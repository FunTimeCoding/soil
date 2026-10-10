package convert

type SlimQueryResult struct {
	Type     string       `json:"type"`
	Vector   []SlimSample `json:"vector,omitempty"`
	Matrix   []SlimStream `json:"matrix,omitempty"`
	Scalar   *SlimPoint   `json:"scalar,omitempty"`
	Warnings []string     `json:"warnings,omitempty"`
	Notices  []string     `json:"notices,omitempty"`
}
