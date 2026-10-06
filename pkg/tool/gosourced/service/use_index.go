package service

import (
	"github.com/funtimecoding/soil/pkg/lint/face"
	"github.com/funtimecoding/soil/pkg/source/index/cache"
)

func (s *Service) UseIndex(directory string) {
	s.workspaces = cache.New(directory, face.Kind())
}
