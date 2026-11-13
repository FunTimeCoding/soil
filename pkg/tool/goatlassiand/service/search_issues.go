package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) SearchIssues(
	query string,
	limit *int,
) ([]*issue.Issue, error) {
	return peer.Guarded(
		context.Background(),
		s.jiraGuard,
		func(context.Context) ([]*issue.Issue, error) {
			if limit != nil {
				return s.jira.SearchLimit(*limit, query)
			}

			return s.jira.Search(query)
		},
	)
}
