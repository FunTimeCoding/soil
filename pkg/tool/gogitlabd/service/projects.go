package service

import "github.com/funtimecoding/soil/pkg/gitlab/project"

func (s *Service) Projects(search string) ([]*project.Project, error) {
	if search != "" {
		return s.client.SearchProject(search)
	}

	return s.client.Projects()
}
