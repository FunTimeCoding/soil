package service

import (
	"context"
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
	"github.com/funtimecoding/soil/pkg/peer"
)

func (s *Service) IssueType(i *issue.Issue) (*jira.MetaIssueType, error) {
	return peer.Guarded(
		context.Background(),
		s.jiraGuard,
		func(context.Context) (*jira.MetaIssueType, error) {
			p, e := s.jira.MetaProject(i.Raw.Fields.Project.Key)

			if e != nil {
				return nil, e
			}

			return p.GetIssueTypeWithName(i.Type), nil
		},
	)
}
