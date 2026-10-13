package mock_service

import "github.com/funtimecoding/soil/pkg/inventory/selection"

func (s *Service) Selection() *selection.Selection {
	return s.selection
}
