package service

import "github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"

func (s *Service) Event(
	project string,
	identifier string,
) (*response.Event, error) {
	return s.client.Event(s.organization, project, identifier)
}
