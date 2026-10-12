package service

import "github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"

func (s *Service) Whoami() (*response.User, error) {
	return s.client.Whoami()
}
