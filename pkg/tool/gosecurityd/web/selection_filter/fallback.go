package selection_filter

func (f *Filter) Fallback(
	minimum float64,
	floors []float64,
) *Filter {
	f.fallback = minimum
	f.floors = floors

	return f
}
