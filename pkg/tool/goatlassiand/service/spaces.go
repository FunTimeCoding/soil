package service

import "github.com/funtimecoding/soil/pkg/atlassian/confluence/space"

func (s *Service) Spaces() ([]*space.Space, error) {
	return guarded(s.confluenceGuard, s.confluence.Spaces)
}
