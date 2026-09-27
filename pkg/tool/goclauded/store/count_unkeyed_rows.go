package store

import "github.com/funtimecoding/soil/pkg/tool/goclauded/constant"

func (s *Store) CountUnkeyedRows() map[string]int64 {
	result := map[string]int64{}

	for _, t := range constant.SessionKeyTables {
		result[t] = unkeyedRows(s.database, t)
	}

	return result
}
