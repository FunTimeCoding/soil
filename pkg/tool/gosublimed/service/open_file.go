package service

import (
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/sublime/view"
)

func (s *Service) OpenFile(path string) (*view.View, error) {
	if path == "" {
		return nil, validation.New("file_path is required")
	}

	return s.source.OpenFile(path)
}
