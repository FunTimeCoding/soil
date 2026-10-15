package origin

type Document struct {
	Version string    `json:"Version"`
	Time    string    `json:"Time"`
	Origin  Reference `json:"Origin"`
}
