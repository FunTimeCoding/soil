package service

import "github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"

func (s *Service) IssueTagValues(
	identifier string,
	tag string,
	limit int,
) ([]response.TagValue, error) {
	return s.client.IssueTagValues(s.organization, identifier, tag, limit)
}
