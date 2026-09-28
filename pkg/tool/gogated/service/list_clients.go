package service

import "github.com/funtimecoding/soil/pkg/tool/gogated/model/client"

func (s *Service) ListClients() []*client.Client {
	return s.store.ListClients()
}
