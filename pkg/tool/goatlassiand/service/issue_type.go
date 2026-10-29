package service

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
)

func (s *Service) IssueType(i *issue.Issue) (*jira.MetaIssueType, error) {
	return guarded(
		s.jiraGuard,
		func() (*jira.MetaIssueType, error) {
			p, e := s.jira.MetaProject(i.Raw.Fields.Project.Key)

			if e != nil {
				return nil, e
			}

			return p.GetIssueTypeWithName(i.Type), nil
		},
	)
}
