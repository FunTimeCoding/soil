package service

import (
	"fmt"
	"github.com/andygrunwald/go-jira"
)

func (s *Service) CreateMeta(
	project string,
	issueType string,
) (*jira.MetaIssueType, error) {
	return guarded(
		s.jiraGuard,
		func() (*jira.MetaIssueType, error) {
			p, e := s.jira.MetaProject(project)

			if e != nil {
				return nil, fmt.Errorf("read project metadata: %w", e)
			}

			t, f := s.jira.MetaIssueType(p, issueType)

			if f != nil {
				return nil, fmt.Errorf("read issue type: %w", f)
			}

			return t, nil
		},
	)
}
