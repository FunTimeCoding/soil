package response

type Summary struct {
	Tool    string `json:"tool"`
	Surface string `json:"surface,omitempty"`
	Kind    string `json:"kind,omitempty"`
	Outcome string `json:"outcome,omitempty"`
	Count   int64  `json:"count"`
}
