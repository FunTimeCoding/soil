package facet_count

func New() *Count {
	return &Count{
		Coverage: map[string]int{},
		Score:    map[float64]int{},
		Labels:   map[string]map[string]int{},
	}
}
