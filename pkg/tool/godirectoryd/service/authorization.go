package service

import "github.com/funtimecoding/soil/pkg/web/authorization/client"

func (s *Service) Authorization() *client.Client {
	return s.authorization
}
