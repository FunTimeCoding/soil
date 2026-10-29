package service

import (
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue/option"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/query"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func (s *Service) closedStatus() (string, error) {
	o, e := guarded[*option.Issue](s.jiraGuard, s.jira.IssueOption)

	if e != nil {
		return "", e
	}

	return join.Comma(query.Quote(o.ClosedStatus)), nil
}
