package service

import "github.com/funtimecoding/soil/pkg/atlassian/jira/issue"

func (s *Service) SearchIssues(
	query string,
	limit *int,
) ([]*issue.Issue, error) {
	return guarded(
		s.jiraGuard,
		func() ([]*issue.Issue, error) {
			if limit != nil {
				return s.jira.SearchLimit(*limit, query)
			}

			return s.jira.Search(query)
		},
	)
}
