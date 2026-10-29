package store

import "github.com/funtimecoding/soil/pkg/tool/gogated/model/client"

func (s *Store) ListClients() ([]*client.Client, error) {
	var result []*client.Client

	return result, s.mapper.Order("created_at desc").Find(&result).Error
}
