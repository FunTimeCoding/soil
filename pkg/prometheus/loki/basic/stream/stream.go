package stream

type Stream struct {
	Stream map[string]string `json:"stream"`
	Values [][]string        `json:"values"`
}
