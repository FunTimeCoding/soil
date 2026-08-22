package service

import (
	"fmt"
	"github.com/andygrunwald/go-jira"
)

func (s *Service) CreateMeta(
	project string,
	issueType string,
) (*jira.MetaIssueType, error) {
	p, e := s.jira.MetaProject(project)

	if e != nil {
		return nil, fmt.Errorf("project metadata not found: %w", e)
	}

	t, f := s.jira.MetaIssueType(p, issueType)

	if f != nil {
		return nil, fmt.Errorf("issue type not found: %w", f)
	}

	return t, nil
}
