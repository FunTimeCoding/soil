package service

import (
	"context"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue/option"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/query"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func (s *Service) closedStatus() (string, error) {
	o, e := peer.Guarded(
		context.Background(),
		s.jiraGuard,
		func(context.Context) (*option.Issue, error) {
			return s.jira.IssueOption()
		},
	)

	if e != nil {
		return "", e
	}

	return join.Comma(query.Quote(o.ClosedStatus)), nil
}
