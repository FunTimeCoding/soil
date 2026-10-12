package service

import "github.com/funtimecoding/soil/pkg/gitlab/file"

func (s *Service) File(
	project int64,
	branch string,
	name string,
) (*file.File, error) {
	return s.client.File(project, branch, name)
}
