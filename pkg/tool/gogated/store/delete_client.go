package store

import "github.com/funtimecoding/soil/pkg/tool/gogated/model/client"

func (s *Store) DeleteClient(identifier string) error {
	return s.mapper.Where(
		"identifier = ?",
		identifier,
	).Delete(client.Stub()).Error
}
