package registry

import "github.com/funtimecoding/soil/pkg/web/palette"

func (r *Registry) Search(query string) []palette.Result {
	if query == "" {
		results := make([]palette.Result, len(r.commands))

		for i, c := range r.commands {
			results[i] = palette.Result{Command: c}
		}

		return results
	}

	var results []palette.Result

	for _, c := range r.commands {
		score, positions := palette.Match(query, c.Label)

		if score < 0 {
			continue
		}

		results = append(
			results,
			palette.Result{Command: c, Score: score, Positions: positions},
		)
	}

	palette.SortResults(results)

	return results
}
