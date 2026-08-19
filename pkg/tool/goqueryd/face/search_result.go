package face

type SearchResult struct {
	Path     string
	Body     string
	Score    float64
	Metadata map[string][]string
}
