package store

import "github.com/funtimecoding/soil/pkg/tool/gogated/model/client"

func (s *Store) ListClients() []*client.Client {
	var result []*client.Client
	s.mapper.Order("created_at desc").Find(&result)

	return result
}
