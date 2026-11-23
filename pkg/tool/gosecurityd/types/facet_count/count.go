package facet_count

type Count struct {
	Coverage map[string]int
	Score    map[float64]int
	Labels   map[string]map[string]int
}
