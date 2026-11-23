package system

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/system_result"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/types/system_selection"
)

func (s *Service) Selected(
	selection *system_selection.Selection,
) ([]*system_result.Detail, error) {
	if e := selection.Check(); e != nil {
		return nil, e
	}

	details, e := s.Details()

	if e != nil {
		return nil, e
	}

	result := make([]*system_result.Detail, 0, len(details))

	for _, d := range details {
		if selection.Matches(d.Summary) {
			result = append(result, d)
		}
	}

	return result, nil
}
