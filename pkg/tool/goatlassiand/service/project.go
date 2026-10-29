package service

import "github.com/andygrunwald/go-jira"

func (s *Service) Project(key string) (*jira.Project, error) {
	return guarded(
		s.jiraGuard,
		func() (*jira.Project, error) { return s.jira.Project(key) },
	)
}
