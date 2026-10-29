package service

import (
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
)

func (s *Service) newest(closed string) ([]*issue.Issue, error) {
	if !s.HasProject() {
		return nil, nil
	}

	return s.SearchIssues(
		NewestQuery(s.project, closed),
		new(constant.NewestLimit),
	)
}
