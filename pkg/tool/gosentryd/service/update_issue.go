package service

import "github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"

func (s *Service) UpdateIssue(
	identifier string,
	status string,
	assignedTo string,
) (*response.Issue, error) {
	return s.client.UpdateIssue(s.organization, identifier, status, assignedTo)
}
