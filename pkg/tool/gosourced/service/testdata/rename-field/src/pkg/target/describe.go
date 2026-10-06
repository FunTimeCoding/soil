package target

func (s *Store) Describe() string {
	return s.api + s.Name
}
