package origin

type Document struct {
	Time   string    `json:"Time"`
	Origin Reference `json:"Origin"`
}
