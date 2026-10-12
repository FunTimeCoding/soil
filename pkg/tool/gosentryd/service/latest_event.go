package service

import "github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"

func (s *Service) LatestEvent(issue string) (*response.Event, error) {
	return s.client.LatestEvent(s.organization, issue)
}
