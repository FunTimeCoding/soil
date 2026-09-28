package service

import "github.com/funtimecoding/soil/pkg/tool/gogated/face"

func (s *Service) WithDirectory(d face.Directory) *Service {
	s.directory = d

	return s
}
