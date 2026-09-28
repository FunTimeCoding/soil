package store

import "github.com/funtimecoding/soil/pkg/tool/gogated/model/client"

func (s *Store) CreateClient(c *client.Client) error {
	return s.mapper.Create(c).Error
}
