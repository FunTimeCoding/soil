package service

import "github.com/andygrunwald/go-jira"

func (s *Service) LinkIssues(
	key string,
	target string,
	linkType string,
) error {
	return s.jira.AddLink(
		&jira.IssueLink{
			Type:         jira.IssueLinkType{Name: linkType},
			OutwardIssue: &jira.Issue{Key: key},
			InwardIssue:  &jira.Issue{Key: target},
		},
	)
}
