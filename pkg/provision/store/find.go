package store

func (s *Store) Find(identifier uint) (*Run, error) {
	var result Run

	return &result, s.mapper.Table(s.tableName).First(&result, identifier).Error
}
