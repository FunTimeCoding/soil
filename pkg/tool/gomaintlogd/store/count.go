package store

import "github.com/funtimecoding/soil/pkg/tool/gomaintlogd/store/entry"

func (s *Store) Count() int {
	var result int64
	s.database.Model(entry.New()).Count(&result)

	return int(result)
}
