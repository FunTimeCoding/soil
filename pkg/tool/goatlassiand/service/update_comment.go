package service

import "github.com/andygrunwald/go-jira"

func (s *Service) UpdateComment(
	key string,
	identifier string,
	body string,
) error {
	return s.jira.UpdateComment(key, &jira.Comment{ID: identifier, Body: body})
}
