package service

import "github.com/andygrunwald/go-jira"

func (s *Service) SearchUsers(query string) ([]jira.User, error) {
	return guarded(
		s.jiraGuard,
		func() ([]jira.User, error) { return s.jira.FindUsers(query) },
	)
}
