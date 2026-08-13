package palette

func (r *Registry) Search(query string) []Result {
	if query == "" {
		results := make([]Result, len(r.commands))

		for i, c := range r.commands {
			results[i] = Result{Command: c}
		}

		return results
	}

	var results []Result

	for _, c := range r.commands {
		score, positions := Match(query, c.Label)

		if score < 0 {
			continue
		}

		results = append(
			results,
			Result{Command: c, Score: score, Positions: positions},
		)
	}

	sortResults(results)

	return results
}
