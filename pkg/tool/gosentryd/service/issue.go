package service

import "github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"

func (s *Service) Issue(identifier string) (*response.Issue, error) {
	return s.client.IssueByIdentifier(s.organization, identifier)
}
