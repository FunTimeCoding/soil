package service

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/sublime/view"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/types/edit"
)

func (s *Service) EditView(change *edit.Edit) (*view.View, error) {
	if change.Old == "" {
		return nil, validation.New("old_string is required")
	}

	return guarded(
		s.guard,
		func() (*view.View, error) {
			return s.client.EditView(
				change.Identifier,
				change.Old,
				change.New,
				change.All,
			)
		},
	)
}
