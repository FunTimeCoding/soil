package store

func (s *Store) Create(label string) error {
	return s.mapper.Create(&Click{Label: label}).Error
}
