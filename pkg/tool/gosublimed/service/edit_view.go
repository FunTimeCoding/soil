package service

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/sublime/view"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/types/edit"
)

func (s *Service) EditView(e *edit.Edit) (*view.View, error) {
	if e.Old == "" {
		return nil, validation.New("old_string is required")
	}

	return s.source.EditView(e.Identifier, e.Old, e.New, e.All)
}
