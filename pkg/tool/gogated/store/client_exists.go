package store

import "github.com/funtimecoding/soil/pkg/tool/gogated/model/client"

func (s *Store) ClientExists(identifier string) bool {
	var count int64
	s.mapper.Model(client.Stub()).Where("identifier = ?", identifier).Count(
		&count,
	)

	return count > 0
}
