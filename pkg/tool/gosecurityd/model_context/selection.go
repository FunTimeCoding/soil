package model_context

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/model_context/argument"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/system_selection"
)

func selection(
	a argument.Selection,
	minimum float64,
) (*system_selection.Selection, error) {
	result := system_selection.New().
		WithCoverage(a.Show).
		WithSearch(a.Search).
		WithMinimum(minimum)

	for key, value := range a.Labels {
		if key == "" || value == "" {
			return nil, validation.New("label filter malformed: %s", key)
		}

		result = result.WithLabel(key, value)
	}

	return result, nil
}
