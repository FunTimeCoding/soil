package service

import "github.com/andygrunwald/go-jira"

func (s *Service) Projects() (*jira.ProjectList, error) {
	return guarded(s.jiraGuard, s.jira.Projects)
}
