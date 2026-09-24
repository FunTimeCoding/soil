package mutation

func Within(
	mutants []*Mutant,
	start int,
	end int,
) []*Mutant {
	var result []*Mutant

	for _, m := range mutants {
		if m.Line >= start && m.Line <= end {
			result = append(result, m)
		}
	}

	return result
}
