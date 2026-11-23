package convert

import "github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/system_selection"

func Selection(
	show *string,
	labels *[]string,
	search *string,
	score *float64,
	fallback float64,
) (*system_selection.Selection, error) {
	result := system_selection.New().
		WithCoverage(text(show)).
		WithSearch(text(search)).
		WithMinimum(fallback)

	if score != nil {
		result = result.WithMinimum(*score)
	}

	if labels == nil {
		return result, nil
	}

	for _, l := range *labels {
		next, e := result.WithLabelText(l)

		if e != nil {
			return nil, e
		}

		result = next
	}

	return result, nil
}
