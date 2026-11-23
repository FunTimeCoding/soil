package system

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/facet_count"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/system_selection"
)

func (s *Service) Facets(
	selection *system_selection.Selection,
) (*facet_count.Count, error) {
	if e := selection.Check(); e != nil {
		return nil, e
	}

	details, e := s.Details()

	if e != nil {
		return nil, e
	}

	result := facet_count.New()

	for _, d := range details {
		result.Add(selection, d.Summary)
	}

	return result, nil
}
