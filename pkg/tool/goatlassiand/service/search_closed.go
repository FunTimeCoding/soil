package service

import "github.com/funtimecoding/soil/pkg/atlassian/jira/issue"

func (s *Service) searchClosed(
	query string,
	closed string,
) ([]*issue.Issue, error) {
	return guarded(
		s.jiraGuard,
		func() ([]*issue.Issue, error) { return s.jira.Search(query, closed) },
	)
}
