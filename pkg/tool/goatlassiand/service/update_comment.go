package service

import (
	"context"
	"github.com/andygrunwald/go-jira"
)

func (s *Service) UpdateComment(
	key string,
	identifier string,
	body string,
) error {
	return s.jiraGuard.Attempt(
		context.Background(),
		func(context.Context) error {
			return s.jira.UpdateComment(
				key,
				&jira.Comment{ID: identifier, Body: body},
			)
		},
	)
}
