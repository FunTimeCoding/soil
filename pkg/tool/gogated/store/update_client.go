package store

import "github.com/funtimecoding/soil/pkg/tool/gogated/model/client"

func (s *Store) UpdateClient(row *client.Client) error {
	return s.mapper.Save(row).Error
}
