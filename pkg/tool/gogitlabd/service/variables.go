package service

import "github.com/funtimecoding/soil/pkg/gitlab/variable"

func (s *Service) Variables(project int64) ([]*variable.Variable, error) {
	return s.client.Variables(project)
}
