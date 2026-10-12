package service

import "github.com/funtimecoding/soil/pkg/gitlab/project"

func (s *Service) MustProjects() []*project.Project {
	return s.client.MustProjects()
}
