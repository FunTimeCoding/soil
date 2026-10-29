package service

import "github.com/andygrunwald/go-jira"

func (s *Service) Transitions(key string) ([]jira.Transition, error) {
	return guarded(
		s.jiraGuard,
		func() ([]jira.Transition, error) { return s.jira.Transitions(key) },
	)
}
