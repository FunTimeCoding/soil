package service

import "github.com/funtimecoding/soil/pkg/gitlab/project"

func (s *Service) Project(identifier int64) (*project.Project, error) {
	return s.client.Project(identifier)
}
