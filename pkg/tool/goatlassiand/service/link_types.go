package service

import "github.com/funtimecoding/soil/pkg/atlassian/jira/link_type"

func (s *Service) LinkTypes() ([]link_type.Type, error) {
	return guarded(s.jiraGuard, s.jira.LinkTypes)
}
