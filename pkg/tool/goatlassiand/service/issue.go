package service

import "github.com/funtimecoding/soil/pkg/atlassian/jira/issue"

func (s *Service) Issue(key string) (*issue.Issue, error) {
	return guarded(
		s.jiraGuard,
		func() (*issue.Issue, error) { return s.jira.Issue(key) },
	)
}
