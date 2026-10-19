package mock_service

import "github.com/funtimecoding/soil/pkg/generative/model_context/selection"

func (s *Service) Selection() *selection.Selection {
	return s.selection
}
