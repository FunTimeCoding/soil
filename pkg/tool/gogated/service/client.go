package service

import "github.com/funtimecoding/soil/pkg/tool/gogated/model/client"

func (s *Service) Client(identifier string) (*client.Client, error) {
	return s.store.ClientRow(identifier)
}
