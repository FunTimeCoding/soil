package store

func (s *Store) CountUnkeyedRows() map[string]int64 {
	result := map[string]int64{}

	for _, t := range sessionKeyTables() {
		result[t] = unkeyedRows(s.database, t)
	}

	return result
}
